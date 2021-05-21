package database

import (
	"fmt"
	"log"
	"strconv"

	//"log"
	"path/filepath"
	"testing"

	"github.com/mitchellh/go-homedir"
)

// test boti() and itob()

var home, _ = homedir.Dir()


var path = filepath.Join(home, "test.db")
func TestInit(t *testing.T) {
	
	Init(path)
	if path != filepath.Join(home,"test.db"){
		fmt.Print("Enable to connect to ...  ",path)
		fmt.Println()
	}else{
		fmt.Print("\nConnecting to ",path)
		fmt.Print("\n\n")
	}
}

func TestAddTask(t *testing.T){
	
	_, record := NewTask("New Record")
	
	if record ==nil{
		fmt.Print("Added new record.\n")
	}else{
			t.Error("Error ")
		}
		fmt.Print("\n**********************************************\n\n")
}


func TestViewTasks(t *testing.T) {
	activities, err := ViewAllTasks()
	for num, tasks := range activities{
		fmt.Printf("%d . %s (%d) \n",num+1, tasks.Task,tasks.Num)
	} 
	fmt.Print("\n**********************************************\n\n")
	if err != nil {
		log.Println("Get all failed")
	}
}

func TestDelTask(t *testing.T) {
args := []string{"1","2"}
	activities, _ := ViewAllTasks()
		var nums []int
		for _, num := range args{
			result, err := strconv.Atoi(num)
			
			if err != nil{
				fmt.Println("\nSome error occured while parsing this:-",num)
				
			}else{
				nums = append(nums, result)
			}
		}
		fmt.Printf("Successfully completed tasks-\n\n")
		fmt.Print("**********************************************\n")

		for _, num := range nums{
			
			if num <= 0 || num > len(activities){	
				fmt.Println()
				fmt.Printf("Wrong task numbers:- %d",num)
				
				continue
			}
			
			activity := activities[num-1]
			
			err := DelTask(2)
			
			if err == nil{
				
				fmt.Printf("%d. %s\n",num+1,activity.Task)
				
				
			}else{
				fmt.Printf("%s enable to done",activity.Task)
			}
		}
		fmt.Print("**********************************************\n")
		fmt.Print("\n")


}
