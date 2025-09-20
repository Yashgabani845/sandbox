#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_WORDS 1000
#define WORD_LEN 20

int wordExists(char words[][WORD_LEN], int size, char *word) {
    for (int i = 0; i < size; i++)
        if (strcmp(words[i], word) == 0) return 1;
    return 0;
}

int ladderLength(char *beginWord, char *endWord, char words[][WORD_LEN], int wordCount) {
    int queue[MAX_WORDS], dist[MAX_WORDS];
    int front = 0, rear = 0;

    queue[rear] = -1;  // beginWord is virtual, index -1
    dist[rear++] = 1;

    int visited[MAX_WORDS] = {0};

    while (front < rear) {
        int u = queue[front];
        int d = dist[front++];
        char *current;
        char tmp[WORD_LEN];
        if (u == -1) current = beginWord;
        else current = words[u];

        if (strcmp(current, endWord) == 0) return d;

        for (int i = 0; i < wordCount; i++) {
            if (!visited[i]) {
                int diff = 0;
                for (int j = 0; current[j]; j++)
                    if (current[j] != words[i][j]) diff++;
                if (diff == 1) {
                    visited[i] = 1;
                    queue[rear] = i;
                    dist[rear++] = d + 1;
                }
            }
        }
    }
    return 0;
}

int main() {
    char beginWord[WORD_LEN] = "hit";
    char endWord[WORD_LEN] = "cog";
    char wordList[MAX_WORDS][WORD_LEN] = {"hot","dot","dog","lot","log","cog"};
    int wordCount = 6;

    int ans = ladderLength(beginWord, endWord, wordList, wordCount);
    printf("%d\n", ans);  // Expected output: 5
    return 0;
}