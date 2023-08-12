package db

import (
	"errors"
	"time"
)

// ToDoItem is the struct that represents a single ToDo item
type voterPoll struct {
	PollID   uint
	VoteDate time.Time
}

// VoterList is a type alias for a map of Voters.  The key
// will be the ToDoItem.Id and the value will be the ToDoItem
// type DbMap map[int]ToDoItem

type Voter struct {
	VoterId     uint
	FirstName   string
	LastName    string
	VoteHistory []voterPoll
}

type VoterList struct {
	Voters map[uint]Voter // A map of VoterIDs as keys and Voter structs as values
}

//------------------------------------------------------------
// THESE ARE THE PUBLIC FUNCTIONS THAT SUPPORT OUR TODO APP
//------------------------------------------------------------

func NewVoter(id uint, fn, ln string) *Voter {
	return &Voter{
		FirstName:   fn,
		LastName:    ln,
		VoteHistory: []voterPoll{},
	}
}

func (v *Voter) AddPoll(pollID uint) {
	v.VoteHistory = append(v.VoteHistory, voterPoll{PollID: pollID, VoteDate: time.Now()})
}

func NewVoterList() (*VoterList, error) {

	voterList := &VoterList{
		Voters: make(map[uint]Voter),
	}

	return voterList, nil
}

func (lst *VoterList) AddVoter(voter Voter) error {

	lst.Voters[voter.VoterId] = voter
	return nil
}

func (lst *VoterList) AddVoterById(voterID uint) (Voter, error) {

	newVoter := Voter{
		VoterId: voterID,
	}
	lst.Voters[voterID] = newVoter
	return newVoter, nil
}

func (lst *VoterList) DeleteVoter(id uint) error {

	// we should if item exists before trying to delete it
	// this is a good practice, return an error if the
	// item does not exist

	//Now lets use the built-in go delete() function to remove
	//the item from our map

	delete(lst.Voters, id)
	return nil
}

func (lst *VoterList) DeleteAll() error {
	//To delete everything, we can just create a new map
	//and assign it to our existing map.  The garbage collector
	//will clean up the old map for us
	lst.Voters = make(map[uint]Voter)

	return nil
}

func (lst *VoterList) UpdateVoter(voter Voter) error {

	// Check if item exists before trying to update it
	// this is a good practice, return an error if the
	// item does not exist
	_, ok := lst.Voters[voter.VoterId]
	if !ok {
		return errors.New("item does not exist")
	}

	//Now that we know the item exists, lets update it
	lst.Voters[voter.VoterId] = voter

	return nil
}

/*
Get a single voter resource with voterID=:id including their entire voting history.
POST version adds one to the "database"
*/
func (lst *VoterList) GetSingleVoterResource(id uint) (Voter, error) {

	// Check if item exists before trying to get it
	// this is a good practice, return an error if the
	// item does not exist
	voter, ok := lst.Voters[id]
	if !ok {
		return Voter{}, errors.New("item does not exist")
	}

	return voter, nil
}

/*
Gets JUST the voter history for the voter with VoterID = :id
*/
func (lst *VoterList) GetVoterHistory(id uint) ([]voterPoll, error) {

	voter, ok := lst.Voters[id]
	if !ok {
		return []voterPoll{}, errors.New("item does not exist")
	}

	return voter.VoteHistory, nil

}

/*
Get all voter resources including all voter history for each voter (note we will
discuss the concept of "paging" later, for now you can ignore)
*/
func (lst *VoterList) GetAllVoters() ([]Voter, error) {

	//Now that we have the DB loaded, lets crate a slice
	var voterList []Voter

	//Now lets iterate over our map and add each item to our slice
	for _, item := range lst.Voters {
		voterList = append(voterList, item)
	}

	//Now that we have all of our items in a slice, return it
	return voterList, nil
}

/*
Gets JUST the single voter poll data with PollID = :id and VoterID = :id.
*/
func (lst *VoterList) GetVoterPollData(voterId uint, pollId uint) (*voterPoll, error) {

	currentVoter, ok := lst.Voters[voterId]
	if !ok {
		return nil, errors.New("item does not exist")
	}

	for j := 0; j < len(currentVoter.VoteHistory); j++ {
		currentPoll := currentVoter.VoteHistory[j]
		if currentPoll.PollID == pollId {
			return &currentPoll, nil
		}
	}

	return nil, errors.New("item does not exist")

}

func (lst *VoterList) AddVoterPollData(voterId uint, pollId uint) error {

	currentVoter, ok := lst.Voters[voterId]
	if !ok {
		newVoter := Voter{
			VoterId: voterId,
		}
		// lst.Voters[voterId] = newVoter

		newPoll := voterPoll{
			PollID:   pollId,
			VoteDate: time.Now(),
		}
		newVoter.VoteHistory = append(newVoter.VoteHistory, newPoll)
		lst.Voters[voterId] = newVoter
		return nil
	}

	newPoll := voterPoll{
		PollID:   pollId,
		VoteDate: time.Now(),
	}
	currentVoter.VoteHistory = append(currentVoter.VoteHistory, newPoll)

	lst.Voters[voterId] = currentVoter
	return nil
}

func (lst *VoterList) DeletePoll(voterId uint, pollId uint) error {

	currentVoter, ok := lst.Voters[voterId]
	if !ok {
		return errors.New("item does not exist")
	}

	index := 0
	for j := 0; j < len(currentVoter.VoteHistory); j++ {
		currentPoll := currentVoter.VoteHistory[j]
		if currentPoll.PollID == pollId {
			index = j
		}
	}

	currentVoter.VoteHistory = append(currentVoter.VoteHistory[:index], currentVoter.VoteHistory[index+1:]...)

	lst.Voters[voterId] = currentVoter
	return nil
}
