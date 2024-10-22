# AI-Learn: Intelligent Learning Assistant

AI-Learn is a command-line tool designed to enhance your learning experience through AI-powered spaced repetition and dynamic problem generation. It helps you efficiently learn and retain knowledge on various topics.

## Key Features

- **Topic Management**: Organize your learning materials into topics.
- **Item Creation**: Add specific items to learn within each topic.
- **AI-Powered Learning**: Generate problems and questions tailored to your learning needs.
- **Spaced Repetition**: Optimize your review schedule for better long-term retention.
- **Progress Tracking**: Monitor your learning progress over time.

## Basic Usage

### Initial Setup

1. Initialize AI-Learn:
   ```
   ai-learn init
   ```

2. Set up your OpenAI API key:
   ```
   ai-learn settings set OPENAI_API_KEY YOUR_API_KEY_HERE
   ```

### Managing Topics and Items

1. Add a new topic:
   ```
   ai-learn topic add DSA
   ```

2. Add items to a topic:
   ```
   ai-learn item add 1 "Two pointer"
   ai-learn item add 1 "Sliding window"
   ```

### Learning and Review

1. Start a learning session for a topic:
   ```
   ai-learn topic learn 1
   ```

2. Provide feedback on an item after review:
   ```
   ai-learn item feedback 2 easy
   ```

## How It Works

1. **Topic Creation**: You create topics for subjects you want to learn.
2. **Item Addition**: Within each topic, you add specific items or concepts to study.
3. **Learning Sessions**: When you start a learning session, AI-Learn generates problems or questions based on your items.
4. **Spaced Repetition**: The system schedules review sessions optimally based on your performance and feedback.
5. **Progress Tracking**: Your learning progress is tracked, allowing you to see your improvement over time.

AI-Learn uses advanced AI models to generate relevant problems and adapt to your learning pace, making your study sessions more effective and engaging.

Start your AI-powered learning journey with AI-Learn today!
