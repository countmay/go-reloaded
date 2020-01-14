package main

import (
	"fmt"
	"strings"

	student ".."
)

func main() {
	str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(student.Split(str, "|=choumi=|"))
	fmt.Println(strings.Split(str, "|=choumi=|"))

	str = "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(student.Split(str, "|=choumi=|"))
	fmt.Println(strings.Split(str, "|=choumi=|"))
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(student.Split(str, "AFJ"))
	fmt.Println(strings.Split(str, "AFJ"))

	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(student.Split(str, "<<==123==>>"))
	fmt.Println(strings.Split(str, "<<==123==>>"))

	str = "HeyyouGoHAHAH"
	fmt.Println(student.Split(str, "HAHA"))
	fmt.Println(strings.Split(str, "HAHA"))
	str = "HelloHAhowHAareHAyou?"
	fmt.Println(student.Split(str, "HA"))
	fmt.Println(strings.Split(str, "HA"))
	str = "HAHAhelloHAHAHEHE"
	fmt.Println(student.Split(str, "HAHA"))
	fmt.Println(strings.Split(str, "HAHA"))
	str = "HAHAHAHAHelloHA"
	fmt.Println(student.Split(str, "HA"))
	fmt.Println(strings.Split(str, "HA"))
	str = "HAHelloHAHAHAhowHAareHAyou?HA"
	fmt.Println(student.Split(str, "HA"))
	fmt.Println(strings.Split(str, "HA"))
	str = "HAHAHAHAHAHAHelloHAHAHAHAHAHAHAhowHAareHAyou?HAHAHAHAHAHA"
	fmt.Println(student.Split(str, "HAHA"))
	fmt.Println(strings.Split(str, "HAHA"))


}
