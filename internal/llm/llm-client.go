package llm

type Client interface {
	GenerateProblem(instructions, topicName, itemName string, generatedProblems []string) (string, error)
}
