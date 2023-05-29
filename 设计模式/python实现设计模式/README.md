### 常用的设计模式学习记录
```text
├─代理模式
│      proxy.py
│      test.txt
│
├─单例模式
│      singleton.py
│
├─外观模式
│      facade.py
│
├─工厂模式
│      abstract_factory.py
│      factory.py
│      factory_method.py
│
├─建造者模式
│      builder.py
│
├─桥模式
│      bridge.py
│
├─模板方法模式
│      template_method.py
│
├─策略模式
│      strategy.py
│
├─组合模式
│      combination.py
│
├─观察者模式
│      observer_pattern.py
│
├─责任链模式
│      chain_responsibility.py
│
└─适配器模式
        adapter.py
```

#### 常见概念
1. 设计模式
    - 设计模式是对软件设计中普遍存在或反复出向的各种问题所提出的解决方案。每一个设计模式系统地被命名、解释和评价了面向对象系统中一个重要和重复出现的设计。

2. 设计模式的分类
    - 创建型模式：工厂方法模式、抽象工厂模式、创建者模式、原型模式、单例模式。隐藏底层模块的逻辑，关注怎么创建对象。
    - 结构型模式：适配器模式、桥模式、组合模式、装饰模式、外观模式、享元模式、代理模式。类之间如何协同工作，应该组成什么结构。
    - 行为型模式：解释器模式、责任链模式、命令模式、迭代器模式、中介者模式、备忘录模式、观察者模式、状态模式、策略模式、访问者模式、模板方法模式。关注行为，也就是方法，应该怎样某些行为。

3. 面向对象
    - 设计模式解决的就是面向对象中的问题。需要指导面向对象的三大特性是 封装、继承和多态 ，封装是把数据和方法封装到类中，继承是类之间复用代码，
    - 多态在Python中默认支持的，Python是一种多态的语言。

4. 接口
    - 接口是若干抽象方法的集合。接口的作用是限制实现接口的类必须按照接口给定的调用方式实现这些方法，对高层模块隐藏了类的内部实现。案例如下：
    ```python
    #!/usr/bin/env python
    # -*- encoding: utf-8 -*-
    '''
    @File    :   interface.py
    @Time    :   2022/09/12 21:21:05
    @Author  :   scliang 
    @Desc    :   None
    '''

    # here put the import lib
    # 引入抽象类
    from abc import ABCMeta, abstractmethod
    import re

    class Payment(metaclass=ABCMeta):
        # abstract class
        @abstractmethod
        def pay(self,money):
            pass


    # 实现抽象类
    class AliPay(Payment):
        def pay(self,money):
        print("支付宝支付%d" %money)

    class WechatPay(Payment):
        def pay(self, money):
            print("微信支付%d元" % money)

    # 直接实例化一个抽象类会直接报错
    # p = Payment()
    # print("abstract method:",p)
    '''
    Traceback (most recent call last):
    File "F:\GoProjects\src\MyPractiseNotes\设计模式\interface.py", line 33, in <module>
        p = Payment()
    TypeError: Can't instantiate abstract class Payment with abstract method pay
    '''

    # 可以实例化抽象类的实现类
    p = AliPay()
    p.pay(100)

    """
    PS F:\GoProjects\src\MyPractiseNotes\设计模式> python .\interface.py
    支付宝支付100
    """
    ```

#### 面向对象的三大特性
- 继承
- 封装
- 多态

#### 创建型模式
- [工厂方法模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F/factory_method.py)
- [抽象工厂模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E5%B7%A5%E5%8E%82%E6%A8%A1%E5%BC%8F/abstract_factory.py)
- [建造者模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E5%BB%BA%E9%80%A0%E8%80%85%E6%A8%A1%E5%BC%8F/builder.py)
- 原型模式
- [单例模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E5%8D%95%E4%BE%8B%E6%A8%A1%E5%BC%8F/singleton.py)


#### 结构型模式
- [适配器模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E9%80%82%E9%85%8D%E5%99%A8%E6%A8%A1%E5%BC%8F/adapter.py)
- [桥模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E6%A1%A5%E6%A8%A1%E5%BC%8F/bridge.py)
- [组合模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E7%BB%84%E5%90%88%E6%A8%A1%E5%BC%8F/combination.py)
- 装饰模式
- [外观模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E5%A4%96%E8%A7%82%E6%A8%A1%E5%BC%8F/facade.py)
- 享元模式
- [代理模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E4%BB%A3%E7%90%86%E6%A8%A1%E5%BC%8F/proxy.py)

#### 行为型模式
- 解释器模式
- [责任链模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E8%B4%A3%E4%BB%BB%E9%93%BE%E6%A8%A1%E5%BC%8F/chain_responsibility.py)
- 命令模式
- 迭代器模式
- 中介者模式
- 备忘录模式
- [观察者模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E8%A7%82%E5%AF%9F%E8%80%85%E6%A8%A1%E5%BC%8F/observer_pattern.py)
- 状态模式
- [策略模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E7%AD%96%E7%95%A5%E6%A8%A1%E5%BC%8F/strategy.py)
- 访问者模式
- [模板方法模式](https://github.com/coderitx/study-note/blob/master/%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/python%E5%AE%9E%E7%8E%B0%E8%AE%BE%E8%AE%A1%E6%A8%A1%E5%BC%8F/%E6%A8%A1%E6%9D%BF%E6%96%B9%E6%B3%95%E6%A8%A1%E5%BC%8F/template_method.py)
