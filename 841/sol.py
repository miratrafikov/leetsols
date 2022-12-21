class Solution:
    rooms: List[List[int]]
    visited: List[bool]

    def canVisitAllRooms(self, rooms: List[List[int]]) -> bool:
        n = len(rooms)
        self.rooms = rooms
        self.visited = [False] * n
        self.visit_room(0)
        
        return not (False in self.visited)

    def visit_room(self, room_number: int):
        self.visited[room_number] = True
        for room_number_to_visit in self.rooms[room_number]:
            if not self.visited[room_number_to_visit]:
                self.visit_room(room_number_to_visit)
