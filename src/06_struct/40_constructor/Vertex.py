# VsCodeなら右クリックで'Run Python File in Terminal'で動きます
# pythonでいうところのコンストラクタをgoでどのように書くかの説明のために引用されています。
class Vertex(object):
    def __init__(self, x, y):
        self._x = y
        self._y = y

    def area(self):
        return self._x * self._y

    def scale(self, i):
        self._x = self._x * i
        self._y = self._y * i

v = Vertex(3, 4)
v.scale(10)
print(v.area())


