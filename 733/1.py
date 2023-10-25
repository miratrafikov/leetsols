from collections import deque

class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, color: int) -> List[List[int]]:
        if image[sr][sc] == color:
            return image
        # visited = [[False]*len(image)]*len(image[0])
        visited = [[False] * len(image[0]) for _ in range(len(image))]
        q = deque()
        q.append([sr,sc])
        visited[sr][sc] = True
        while len(q) > 0:
            point = q.pop()
            print(point)
            point_color = image[point[0]][point[1]]
            if point[0]-1 >= 0 and not visited[point[0]-1][point[1]] and image[point[0]-1][point[1]] == point_color:
                q.append([point[0]-1, point[1]])
                visited[point[0]-1][point[1]] = True
            if point[0]+1 < len(image) and not visited[point[0]+1][point[1]] and image[point[0]+1][point[1]] == point_color:
                q.append([point[0]+1, point[1]])
                visited[point[0]+1][point[1]] = True
            if point[1]-1 >= 0 and not visited[point[0]][point[1]-1] and image[point[0]][point[1]-1] == point_color:
                q.append([point[0], point[1]-1])
                visited[point[0]][point[1]-1] = True
            if point[1]+1 < len(image[0]) and not visited[point[0]][point[1]+1] and image[point[0]][point[1]+1] == point_color:
                q.append([point[0], point[1]+1])
                visited[point[0]][point[1]+1] = True
            image[point[0]][point[1]] = color
        return image
