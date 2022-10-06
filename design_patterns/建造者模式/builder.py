#!/usr/bin/env python
# -*- encoding: utf-8 -*-
'''
@File    :   builder.py
@Time    :   2022/09/13 10:56:48
@Author  :   scliang 
@Desc    :   建造者模式
'''

'''
讲一个复杂对象与它的表示分离，使得同样的构建过程可以创建不同的表示.
建造者模式与抽象工厂模式相似，也用来创建复杂对象，主要区别在于建造者模式着重一步步构造一个复杂的对象，而抽象工厂模式着重于多个系列的产品对象.
角色：
    抽象建造者
    具体建造者
    指挥者 ： 用来控制组装顺序
    产品
优点： 
    1. 隐藏了对象创建的细节
    2. 隐藏一个产品的内部结构与装配过程.
    3. 将构造代码与表示代码分开.
    4. 可以对构造过程进行更精细的控制.
'''

# here put the import lib
from abc import ABCMeta,abstractmethod

class Player():
    def __init__(self,face=None,body=None,leg=None) -> None:
        self.face = face
        self.body = body
        self.leg = leg

    def __str__(self) -> str:
        return "face: %s body: %s leg: %s" %(self.face,self.body,self.leg)

# 抽象模式
class PlayerBuilder(metaclass=ABCMeta):
    @abstractmethod
    def build_face(self):
        pass

    @abstractmethod
    def build_body(self):
        pass

    @abstractmethod
    def build_leg(self):
        pass

# 具体模式实现
class SexyGirlBuilder(PlayerBuilder):
    def __init__(self) -> None:
        self.player = Player()

    def build_face(self):
        self.player.face = "漂亮脸蛋."
    
    def build_body(self):
        self.player.body = "完美身材."

    def build_leg(self):
        self.player.leg = "大长腿."

class MonsterBuilder(PlayerBuilder):
    def __init__(self) -> None:
        self.player = Player()

    def build_face(self):
        self.player.face = "怪兽脸."

    def build_body(self):
        self.player.body = "怪兽身材."
    
    def build_leg(self):
        self.player.leg = "怪兽大粗腿."

# 控制构造顺序
class PlayerDirector():
    def build_player(self,builder):
        builder.build_body()
        builder.build_leg()
        builder.build_face()
        return builder.player


# client
girl = SexyGirlBuilder()
builder = PlayerDirector()
p = builder.build_player(girl)
print(p) 