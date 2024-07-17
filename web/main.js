const form = document.querySelector("#scheduleForm");

const submitForm = () => {
    var formData = new FormData(form);

    var data = {};
    formData.forEach(function (value, key) {
        data[key] = value;
    });
    var jsonData = JSON.stringify(data);
    
    // Отправляем JSON на сервер
    fetch('http://localhost:8085/create-schedule', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: jsonData
    })
        .then(response => response.json())
        .then(parsedData => {
            console.log('Успешно:', parsedData);
            // Преобразуем JSON в объект JavaScript

            // Выводим данные в консоль
            console.log(parsedData);

            // Пример вывода данных в HTML
            const weekScheduleElement = document.getElementById('week-schedule');

            // Перебираем расписание недели
            parsedData.week_schedule.forEach(daySchedule => {
                const dayName = daySchedule.day_name;
                const dayClasses = daySchedule.day_classes;

                // Создаем элемент для дня
                const dayElement = document.createElement('div');
                dayElement.classList.add('col-2')
                dayElement.classList.add('bg-light')
                dayElement.classList.add('m-1')
               
                dayElement.textContent = `День: ${dayName}`;

                // Добавляем элемент для дня в контейнер
                weekScheduleElement.appendChild(dayElement);

                // Перебираем классы в дне
                dayClasses.forEach(classSchedule => {
                    const className = classSchedule.class_name;
                    const classSubjects = classSchedule.class_subjects;

                    // Создаем элемент для класса
                    const classElement = document.createElement('div');
                    classElement.textContent = `Класс: ${className}`;

                    // Создаем элемент для предметов в классе
                    const subjectsElement = document.createElement('ul');

                    // Перебираем предметы в классе
                    classSubjects.forEach(subject => {
                        const subjectName = subject.subject_name;
                        const subjectQueue = subject.subject_queue;

                        // Создаем элемент для предмета
                        const subjectItem = document.createElement('li');
                        subjectItem.style.cssText = 'list-style-type: none'


                        subjectItem.textContent = `${subjectName}, Время: ${subjectQueue}`;

                        // Добавляем элемент для предмета в список предметов
                        subjectsElement.appendChild(subjectItem);
                    });

                    // Добавляем список предметов в элемент для класса
                    classElement.appendChild(subjectsElement);

                    // Добавляем элемент для класса в контейнер
                    dayElement.appendChild(classElement);
                });
            });
        });
}