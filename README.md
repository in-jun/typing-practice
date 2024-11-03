# Typing Practice

Programming keyword-based typing practice tool. Improves typing speed while familiarizing with common code syntax.

## Features

- Real-time WPM/CPM tracking
- Accuracy percentage calculation
- Color-coded feedback
- Progress bar visualization
- Programming language keywords (SQL, Java, Python, C++, etc.)

## Requirements

- C++11 or later
- Terminal with ANSI color support

## Build

```bash
g++ -std=c++11 main.cpp -o typing-practice
./typing-practice
```

## Usage

Type the displayed words as fast and accurately as possible. Press Enter after each word. Statistics update in real-time.

## Interface

```
TYPING PRACTICE

     0 WPM     0 CPM  100%  00:00  ──────────────────── (0%)

  drop  query  where  some

❯
```

Statistics are shown at the top: words per minute, characters per minute, accuracy percentage, elapsed time, and progress bar.
