# 原型模式
## 目的
能够复制已有对象， 而又无需使代码依赖它们所属的类。
## 应用场景
1. 如果需要复制一些对象， 同时又希望代码独立于这些对象所属的具体类， 可以使用原型模式。
2. 如果子类的区别仅在于其对象的初始化方式， 那么可以使用该模式来减少子类的数量。 别人创建这些子类的目的可能是为了创建特定类型的对象。