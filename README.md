

```markdown
# 🧠 GAI Chatbot  

A command-line chatbot powered by Google's Gemini AI, written in Go. It provides a sleek "hacker-style" interface with color effects and animations for a better user experience.  

## 🚀 Features  
- AI-powered responses using Google Gemini API  
- Hacker-style UI with animations and colors  
- Typing effect for a more interactive chat  
- Error handling and validation  

## 🛠 Installation  

### 1️⃣ Clone the Repository  
```sh
git clone https://github.com/temi3500/chat-bot.git
cd chat-bot
```

### 2️⃣ Install Dependencies  
```sh
go mod tidy
```

### 3️⃣ Set Up Environment Variables  
Create a `.env` file (or copy `.env.example`):  
```sh
cp .env.example .env
```
Edit `.env` and add your **Gemini API key**:  
```
GEMINI_API_KEY=your-api-key-here
```

### 4️⃣ Run the Chatbot  
```sh
go run main.go chat
```

## 📜 Usage  
Once started, type your messages and receive AI responses. To exit, type:  
```
exit
```

## 🔧 Troubleshooting  
- If dependencies are missing, run:  
  ```sh
  go mod tidy
  ```
- If you see an API error, ensure your **Gemini API key** is correct and active.  

## 📜 License  
MIT License - Use freely, but give credit.  

## ✨ Author  
Created by **Tameem Ahmad Shahzad**  

---
🔥 **Enjoy chatting with AI in style!**  
```

Let me know if you want any modifications! 🚀