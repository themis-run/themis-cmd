package cmd

import "fmt"

func init() {
	Register("nodeinfo", NodeInfo)
}

type NodeResult struct {
	leaderName string
	nodes      []Node
}

func NewNodeResult(leaderName string, nodes []Node) *NodeResult {
	return &NodeResult{
		leaderName: leaderName,
		nodes:      nodes,
	}
}

func (n *NodeResult) String() string {
	str := fmt.Sprintf("leader_name: %s\n", n.leaderName)
	for _, v := range n.nodes {
		str += "{"
		str += fmt.Sprintf(
			"\tname: %s\n\taddress: %s\n\traft_address: %s\n\tterm: %d\n\trole: %s\n\tlog_term: %d\n\tlog_index: %d\n",
			v.Name, v.Address, v.RaftAddress, v.Term, v.Role, v.LogTerm, v.LogIndex)
		str += "}"
	}

	return str
}

func NodeInfo(args ...string) Result {
	if len(args) != 0 {
		return NewErrorResult(errArgsEmpty)
	}

	c := GetThemisClient()
	servers := c.Info().Servers
	leader := c.Info().LeaderName
	nodeList := make([]Node, 0)

	for s := range servers {
		info, err := c.NodeInfo(s)
		if err != nil {
			continue
		}

		nodeList = append(nodeList, Node{
			Name:        info.Name,
			Address:     info.Address,
			RaftAddress: info.RaftAddress,
			Term:        info.Term,
			Role:        info.Role,
			LogTerm:     info.LogTerm,
			LogIndex:    info.LogIndex,
		})
	}

	return NewNodeResult(leader, nodeList)
}

type Node struct {
	Name        string
	Address     string
	RaftAddress string
	Term        int32
	Role        string
	LogTerm     int32
	LogIndex    int32
}
