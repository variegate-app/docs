# Active Record (Активная запись)

## [<<< ---](../sourcedata.md)

![image.png](Active%20Record%20(%D0%90%D0%BA%D1%82%D0%B8%D0%B2%D0%BD%D0%B0%D1%8F%20%D0%B7%D0%B0%D0%BF%D0%B8%D1%81%D1%8C)%203ebc719527034d02a35feed0df412e82/image.png)

Один объект управляет и данными, и поведением. Большинство этих данных постоянны и их надо хранить в БД. Этот паттерн использует подход - хранение логики доступа к данным в объекте сущности.

Объект является "обёрткой" одной строки из БД или представления, включает в себя доступ к БД и логику обращения с данными.

Пример: объект "Person" содержит данные об одной персоне и методы: добавить, обновить или удалить. По сути паттерн продвигает идею когда модель и логика сосредоточены в 1 классе.

рассмотрим диаграмму класса "Person":