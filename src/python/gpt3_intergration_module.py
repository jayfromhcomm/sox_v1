import openai
openai.api_key = "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"
response = openai.Completion.create(
  engine="text-davinci-002",
  prompt="say hey Sox or ask a question: '{}'",
  max_tokens=1500
)
