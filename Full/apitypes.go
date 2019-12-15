package main

// AppID is a generic string identifier used throughout the program
type AppID string

// VoterID is an identifier for a Voter
type VoterID AppID

// CandidateID is an identifier for a Candidate
type CandidateID AppID

// ElectionID is an identifier for an Election
type ElectionID AppID

// TicketID is an identifier for a Ticket
type TicketID AppID

// TicketEntryID is an identifier for a TicketEntry
type TicketEntryID AppID

// ElectionPositionID is an identifier for an ElectionPosition
type ElectionPositionID AppID

// VotePriority is the priority for a vote. This is only used in non FPTP systems
type VotePriority int

// ElectionInfo is a brief description of a given election without including the whole structure
type ElectionInfo struct {
	ID          ElectionID   `json:"id"`
	DisplayName string       `json:"displayName"`
	Term        string       `json:"term"`
	Type        ElectionType `json:"type"`
}

// Election is a given election
type Election struct {
	ID            ElectionID      `json:"id"`
	DisplayName   string          `json:"displayName"`
	Term          string          `json:"term"`
	Type          ElectionType    `json:"type"`
	TicketEntries []TicketEntry   `json:"ballotEntries"`
	Options       ElectionOptions `json:"options,omitempty"`
}

// ElectionOptions are options that apply to a given Election
type ElectionOptions struct {
	CanHaveMultiTicket         bool `json:"canHaveMultiTicket,omitempty"`
	CandidateCanRunForMultiple bool `json:"candidateCanRunForMultiple,omitempty"`
	CandidateCanVote           bool `json:"candidateCanVote,omitempty"`
	CandidateCanVoteForSelf    bool `json:"candidateCanVoteForSelf,omitempty"`
}

// The ElectionType is an enumeration of different possible election types
type ElectionType int

// These constants represent different election types
const (
	FirstPastThePost ElectionType = iota
	InstantRunoff
	MultiRunoff
)

// TicketEntry is an entry in an election.
// For instance, this would define an entry such as "President", or even
// "President and Vice-President", running on the same ticket.
type TicketEntry struct {
	ID                       TicketEntryID        `json:"id"`
	DisplayName              string               `json:"displayName"`
	AllowedElectionPositions []ElectionPositionID `json:"allowedElectionPositions"`
	Tickets                  []TicketID           `json:"tickets"`
}

// Ticket is a specific ticket associated with a list of candidates
// running for this TicketEntry. For instance, you may have one ElectionPositionEntry
// for a President, and another ElectionPositionEntry for a Vice-President
type Ticket struct {
  ID TicketID                                     `json:"id"`
	ElectionPositionEntries []ElectionPositionEntry `json:"electionPositionEntries"`
	Votes                   []Vote                  `json:"votes"`
}

// ElectionPositionEntry is a single candidate and the office that they are running for.
type ElectionPositionEntry struct {
	CandidateID        CandidateID        `json:"candidateId"`
	ElectionPositionID ElectionPositionID `json:"electionPositionId"`
}

// ElectionPosition is a particular position
type ElectionPosition struct {
	ID          ElectionPositionID `json:"id"`
	DisplayName string             `json:"displayName"`
}

// Vote is a vote that can go toward a particular candidate
type Vote struct {
	VoterID      VoterID      `json:"voterId"`
	TicketID     TicketID     `json:"ticketId"`
	VotePriority VotePriority `json:"votePriority"` // used in rank based voting. for now always 1
}

// Voter is a user that is able to vote
type Voter struct {
	ID          VoterID          `json:"id"`
	Permissions VoterPermissions `json:"permissions"`
}

// VoterPermissions are the permissions granted to a given Voter user
type VoterPermissions struct {
	CanCreateElection bool         `json:"canCreateElection"`
	CanManageElection []ElectionID `json:"canManageElection"`
	CanVote           []AppID      `json:"canVote"`
}

// Candidate is a candidate user that is publicly identified, and is able to run in an election
type Candidate struct {
	ID          CandidateID          `json:"id"`
	DisplayName string               `json:"displayName"`
	Permissions CandidatePermissions `json:"permissions"`
}

//CandidatePermissions are the permissions granted to a given Candidate user
type CandidatePermissions struct {
	CanRun []AppID `json:"canRun"`
}

// Elections is example data for elections
var Elections []Election = []Election{
	{
		ID:          "election1",
		DisplayName: "West Chester University Executive Board",
		Term:        "Spring 2020",
		Type:        FirstPastThePost,
		TicketEntries: []TicketEntry{
			{
				ID:                       "ticketentry1",
				DisplayName:              "Presidential",
				AllowedElectionPositions: []ElectionPositionID{"position1"},
        Tickets: []TicketID{ "ticket1", "ticket2" },
			}, {
				ID:                       "ticketentry2",
				DisplayName:              "Secretorial",
				AllowedElectionPositions: []ElectionPositionID{"position2"},
        Tickets: []TicketID{ "ticket3", "ticket4" },
        },
			},
		},
		Options: ElectionOptions{
			CanHaveMultiTicket:         false,
			CandidateCanRunForMultiple: false,
			CandidateCanVote:           true,
			CandidateCanVoteForSelf:    false,
		},
	},
}

// Positions is example data for positions
var Positions []ElectionPosition = []ElectionPosition{
	{ID: "position1", DisplayName: "President"},
	{ID: "position2", DisplayName: "Secretary"},
}

// Voters is example data for voters
var Voters []Voter = []Voter{
	{ID: "voter1", Permissions: VoterPermissions{CanCreateElection: true, CanManageElection: []ElectionID{"election1"}, CanVote: []AppID{"election1"}}},
	{ID: "voter2", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter3", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter4", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter5", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter6", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter7", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
	{ID: "voter8", Permissions: VoterPermissions{CanCreateElection: false, CanManageElection: []ElectionID{}, CanVote: []AppID{"election1"}}},
}

// Tickets is example data for tickets
var Tickets []Ticket = []Ticket{
  {
    ID: "ticket1",
    ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate1", ElectionPositionID: "position1"}},
    Votes:                   []Vote{
		{VoterID: "voter3", TicketID: "ticket1", VotePriority: 1},
		{VoterID: "voter4", TicketID: "ticket1", VotePriority: 1},
		{VoterID: "voter5", TicketID: "ticket1", VotePriority: 1},
	},
  }, {
    ID: "ticket2",
    ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate2", ElectionPositionID: "position1"}},
    Votes:                   []Vote{
		{VoterID: "voter6", TicketID: "ticket2", VotePriority: 1},
		{VoterID: "voter7", TicketID: "ticket2", VotePriority: 1},
		{VoterID: "voter8", TicketID: "ticket2", VotePriority: 1},
	},
  }, {
    ID: "ticket3",
    ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate3", ElectionPositionID: "position2"}},
    Votes:                   []Vote{
		{VoterID: "voter1", TicketID: "ticket3", VotePriority: 1},
		{VoterID: "voter2", TicketID: "ticket3", VotePriority: 1},
		{VoterID: "voter5", TicketID: "ticket3", VotePriority: 1},
	},
  }, {
    ID: "ticket4",
    ElectionPositionEntries: []ElectionPositionEntry{{CandidateID: "candidate4", ElectionPositionID: "position2"}},
    Votes:                   []Vote{
		{VoterID: "voter6", TicketID: "ticket4", VotePriority: 1},
		{VoterID: "voter7", TicketID: "ticket4", VotePriority: 1},
		{VoterID: "voter8", TicketID: "ticket4", VotePriority: 1},
	},
  },
};


// Candidates is example data for candidates
var Candidates []Candidate = []Candidate{
	{ID: "candidate1", DisplayName: "James Brennen", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
	{ID: "candidate2", DisplayName: "Mike Grimson", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
	{ID: "candidate3", DisplayName: "Alicia Michaels", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
	{ID: "candidate4", DisplayName: "Kelly Zimmerman", Permissions: CandidatePermissions{CanRun: []AppID{"election1"}}},
}
