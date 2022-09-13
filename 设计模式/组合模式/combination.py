#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   combination.py
@Time    :   2022/09/13 12:36:17
@Author  :   scliang 
@Desc    :   组合模式
'''

'''
将对象组合成树形结构以表示“部分-整体”的层次结构，组合模式使的用户对单个对象和组合对象的使用具有一致性.
角色：
    抽象组件
    叶子组件
    复合组件
    客户端


'''


# here put the import lib
from abc import ABCMeta, abstractmethod

# 抽象组件
class Graphic(metaclass=ABCMeta):
    @abstractmethod
    def draw(self):
        pass

# 叶子组件
class Point(Graphic):
    def __init__(self, x, y):
        self.x = x
        self.y = y

    def __str__(self):
        return '点(%s,%s)' % (self.x, self.y)

    def draw(self):
        print(self)

# 叶子组件
class Line(Graphic):
    def __init__(self, p1, p2):
        self.p1 = p1
        self.p2 = p2

    def __str__(self):
        return '线段[(%s,%s)]' % (self.p1, self.p2)

    def draw(self):
        print(self)

# 复合组件
class Picture(Graphic):
    def __init__(self, iterable):
        self.children = []
        for g in iterable:
            self.add(g)

    def add(self, graphic):
        self.children.append(graphic)

    def draw(self):
        for g in self.children:
            g.draw()

# 简单图形
print('------简单图形------')
p = Point(1, 2)
l1 = Line(Point(1, 2), Point(3, 4))
l2 = Line(Point(5, 6), Point(7, 8))
print(p)
print(l1)
print(l2)
print('------复合图形(p,l1,l2)------')
# 复合图形
pic = Picture([p, l1, l2])
pic.draw()