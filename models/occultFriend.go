package models

import (
	"log"
	"math/rand"
	"src/amigOculto/database"
	"strings"
)

type OccultFriend struct {
	ID           int
	participants string
	notDraw      string
}

func CreateNew(participants string) {
	db := database.Conect()

	insertInDB, err := db.Prepare("insert into occultfriend(notdraw, participants) values($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	insertInDB.Exec(participants, participants)
	defer db.Close()
}

func GetRandomNotDraw(idOc string) string {
	db := database.Conect()

	oc := OccultFriend{}
	query := "SELECT notdraw FROM occultfriend WHERE id = $1"

	err := db.QueryRow(query, idOc).Scan(&oc.notDraw)
	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	defer db.Close()
	notDrawArray := strings.Split(oc.notDraw, ",")
	randomIndex := rand.Intn(len(notDrawArray))
	pick := notDrawArray[randomIndex]
	return pick

}

func RemoveNotDraw(idOc string, pick string) {
	db := database.Conect()
	oc := OccultFriend{}
	query := "SELECT id,participants, notdraw FROM occultfriend WHERE id=$1"

	err := db.QueryRow(query, idOc).Scan(&oc.ID, &oc.participants, &oc.notDraw)
	if err != nil {
		log.Fatal("1 - Failed to execute query: ", err)
	}
	notDrawArray := strings.Split(oc.notDraw, ",")
	var index int
	for i, element := range notDrawArray {
		if element == pick {
			index = i
		}
	}
	notDrawArray = remove(notDrawArray, index)
	oc.notDraw = strings.Join(notDrawArray, ",")

	newOccultFriend, err := db.Prepare("update occultfriend set notDraw=$1 where id=$2")
	if err != nil {
		panic(err.Error())
	}
	newOccultFriend.Exec(oc.notDraw, idOc)
	defer db.Close()
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
