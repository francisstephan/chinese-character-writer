package data

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// _ "github.com/mattn/go-sqlite3"

type Zi struct {
	id         int
	pinyin_ton string
	unicode    string
	hanzi      string
	sens       string
}

type Dico []Zi

var nmax int
var quizdico Dico

func litdic(where string) (dic Dico) { // read dictionary with WHERE clause
	dic = make(Dico, 0, 10) // initialize dic with size 0, capacity 10
	db, err := sql.Open("sqlite3", "vol/zidian.db")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer db.Close()

	query := "SELECT id, pinyin_ton, unicode, sens FROM pyhz"
	if where != "" {
		query += " WHERE " + where
	}
	// log.Println("requête = ", query)

	rows, err := db.Query(query)
	if err != nil {
		log.Println("erreur=", err.Error())
		return nil
	}
	var tempZi Zi
	for rows.Next() {
		rows.Scan(&tempZi.id, &tempZi.pinyin_ton, &tempZi.unicode, &tempZi.sens)
		intrune, _ := strconv.ParseInt(tempZi.unicode, 16, 64)
		tempZi.hanzi = string(rune(intrune))
		dic = append(dic, tempZi)
	}
	return dic
}

func printdic(sousdic Dico, py string) string { // prepare browser display of partial dictionary
	if len(sousdic) == 0 {
		return "No result for " + py + " query"
	}
	retour := "<p>Result for query '" + py + "' :</p>" +
		"<table><tr><td> Pinyin </td><td> Unicode </td><td> Character </td><td> Translation</td></tr>"
	for _, zi := range sousdic {
		retour += fmt.Sprintf(
			"<tr><td>%7s</td><td>%4s</td><td>%s</td><td>%s</td></tr>", zi.pinyin_ton, zi.unicode, zi.hanzi, zi.sens)
	}
	return retour + "</table>"
}

func Dicsize() string { // get dictionary size = number of entries
	dic := litdic("")
	if dic != nil {
		return strconv.Itoa(len(dic))
	}
	return "zero"
}

func Listforpy(py string) template.HTML { // cf https://pkg.go.dev/html/template#HTML and https://github.com/gin-gonic/gin/issues/858
	last := py[len(py)-1]
	numeric := "01234"
	var where string
	if !strings.Contains(numeric, string(last)) {
		where = "pinyin_ton='" + py + "0' OR pinyin_ton='" + py + "1' OR pinyin_ton='" + py + "2' OR pinyin_ton='" + py +
			"3' OR pinyin_ton='" + py + "4' ORDER BY pinyin_ton , unicode"
	} else {
		where = "pinyin_ton='" + py + "' ORDER BY unicode"
	}
	dic := litdic(where)
	return template.HTML(printdic(dic, py))
}

func Listforzi(zi string) template.HTML { // cf https://pkg.go.dev/html/template#HTML and https://github.com/gin-gonic/gin/issues/858
	r := []rune(zi)
	where := "unicode='" + fmt.Sprintf("%X", r[0]) + "' ORDER BY pinyin_ton "
	dic := litdic(where)
	return template.HTML(printdic(dic, zi))
}

func GetZiList(py string) template.HTML { // get list of possible zi that may be added by clicking on a button
	last := py[len(py)-1]
	numeric := "01234"
	var where string
	if !strings.Contains(numeric, string(last)) { // = if no tone (0 to 4) is included with the pinyin
		where = "pinyin_ton='" + py + "0' OR pinyin_ton='" + py + "1' OR pinyin_ton='" + py + "2' OR pinyin_ton='" + py + "3' OR pinyin_ton='" + py + "4'"
	} else {
		where = "pinyin_ton='" + py + "'"
	}
	dic := litdic(where)
	if len(dic) == 0 {
		return template.HTML("No result for " + py + "query")
	} else {
		retchain := " Choose character from following list :<br />"
		for _, zi := range dic {
			retchain += "<button class='zilistbutton' onclick='add(\"" + zi.hanzi + "\")'>" + zi.hanzi + "</button>"
		}
		return template.HTML(retchain)
	}
}

func GetQuizZi() template.HTML { // get random chose zi as quiz
	n := rand.Intn(nmax)
	zi := quizdico[n-1]
	r := []rune(zi.hanzi)
	where := "unicode='" + fmt.Sprintf("%X", r[0]) + "' ORDER BY pinyin_ton "
	dic := litdic(where)
	return template.HTML(zi.hanzi + "<button type='button' autofocus id='quizbutton' onclick = 'visible(); this.style.display = \"none\"'>Show result</button></p><div id=\"solution\">" + printdic(dic, zi.hanzi) + "</div>")
	// return Listforzi(zi.hanzi)
}

func init() { // initialize the data package with a test on database availability
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(version)

	rand.Seed(time.Now().UnixNano()) // https://golangdocs.com/generate-random-numbers-in-golang
	quizdico = litdic("")
	nmax = len(quizdico)
}
