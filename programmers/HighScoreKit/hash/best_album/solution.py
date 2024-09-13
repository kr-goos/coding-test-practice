from typing import *

def solution(genres: List[str], plays: List[int]) -> List[int]:
    genre_play_count = {}
    genre_songs = {}

    for i in range(len(genres)):
        genre = genres[i]
        play = plays[i]
        
        if genre in genre_play_count:
            genre_play_count[genre] += play
        else:
            genre_play_count[genre] = play
        
        if genre in genre_songs:
            genre_songs[genre].append((i, play))
        else:
            genre_songs[genre] = [(i, play)]
    
    sorted_genres = sorted(genre_play_count.keys(), key=lambda g: genre_play_count[g], reverse=True)

    answer = []
    for genre in sorted_genres:
        songs = genre_songs[genre]
        sorted_songs = sorted(songs, key=lambda x: (-x[1], x[0]))

        answer.extend([song[0] for song in sorted_songs[:2]])

    return answer




genres = ["classic", "pop", "classic", "classic", "pop"]
plays = [500, 600, 150, 800, 2500]
result = [4, 1, 3, 0]
s = solution(genres, plays)
print('solution return : ', s)
print('solution result : ', result)

if s == result:
    print('test success')
else:
    print('test failure')

