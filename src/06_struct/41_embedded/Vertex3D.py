# VsCodeなら右クリックで'Run Python File in Terminal'で動きます
class Vertex(object):
    def __init__(self, x, y):
        self._x = y
        self._y = y

    def area(self):
        return self._x * self._y

    def scale(self, i):
        self._x = self._x * i
        self._y = self._y * i


# Vertexを継承するVertex3D
class Vertex3D(Vertex):
    def __init__(self, x, y, z):
        super().__init__(x, y)
        self._z = z

    def area_3d(self):
        # 縦×横×高さ
        return self._x * self._y * self._z

    def scale_3d(self, i):
        self._x = self._x * i
        self._y = self._y * i
        self._z = self._z * i

v = Vertex3D(3, 4, 5)
v.scale(10)
print(v.area())
print(v.area_３d())

