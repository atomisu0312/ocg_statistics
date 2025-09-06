import dotenv from 'dotenv';
import * as fs from 'fs';
import * as path from 'path';
dotenv.config();

const apiKey = process.env.OPENAI_API_KEY;

if (!apiKey) {
  console.error('OpenAI APIキーが設定されていません。環境変数 OPENAI_API_KEY を設定してください。');
  process.exit(1);
}

const diffFilePath = process.argv[2];
if (!diffFilePath) {
    console.error('diff file path is not provided.');
    process.exit(1);
}
const outputFilePath = process.argv[3];
if (!outputFilePath) {
    console.error('output file path is not provided.');
    process.exit(1);
}
const pastCommentsFilePath = process.argv[4];
let pastComments = '';
if (pastCommentsFilePath && fs.existsSync(pastCommentsFilePath)) {
    pastComments = fs.readFileSync(pastCommentsFilePath, 'utf-8');
}

const changes = fs.readFileSync(diffFilePath, 'utf-8');

interface ChatCompletionMessage {
  role: 'system' | 'user' | 'assistant';
  content: string;
}

interface ChatCompletionResponse {
  id: string;
  object: string;
  created: number;
  model: string;
  choices: {
    index: number;
    message: ChatCompletionMessage;
    finish_reason: string;
  }[];
  usage: {
    prompt_tokens: number;
    completion_tokens: number;
    total_tokens: number;
  };
}

async function getCodeReview(): Promise<void> {
  try {
    const templatePath = path.resolve(__dirname, 'prompt_template.md');
    const template = fs.readFileSync(templatePath, 'utf-8');

    const parts = template.split('## 差分');
    const systemPrompt = parts[0];
    let userPromptTemplate = '## 差分' + parts[1];

    let userPrompt = userPromptTemplate.replace('${change}', changes);
    if (pastComments) {
        userPrompt += `\n\n## 過去のレビューコメント\n${pastComments}`;
    }

    const response = await fetch(
      'https://api.openai.com/v1/chat/completions',
      {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          'Authorization': `Bearer ${apiKey}`
        },
        body: JSON.stringify({
          model: 'gpt-4o-mini',
          messages: [
            {
              role: 'system',
              content: systemPrompt.trim()
            },
            {
              role: 'user',
              content: userPrompt.trim()
            }
          ],
          temperature: 0.7,
          max_tokens: 1500
        })
      }
    );

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(`APIエラー: ${response.status} ${response.statusText} - ${JSON.stringify(errorData)}`);
    }

    const data = await response.json() as ChatCompletionResponse;

    if (data.choices && data.choices.length > 0) {
      console.log('--- ChatGPTによるコードレビュー ---');
      console.log(data.choices[0].message.content);
      console.log('---------------------------------');
      fs.writeFileSync(outputFilePath, data.choices[0].message.content);
    } else {
      console.log('レビュー結果がありませんでした。');
      fs.writeFileSync(outputFilePath, 'レビュー結果がありませんでした。');
    }

  } catch (error) {
    console.error('リクエスト中にエラーが発生しました:', error);
  }
}

getCodeReview();
