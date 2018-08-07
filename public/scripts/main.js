/*jshint esversion: 6 */

new Vue({
    el: '.container',
    data: {
        tasks: [],
        newTask: { name: '', id: '' }
    },
    // Get data from db, on created
    created() {
        //Получаем данные при создании страницы
        this.getData();
    },
    // Get data from db, on updated
    updated() {
        this.getData();
    },
    methods: {
        //get data from db
        getData() {
            // form /tasks rote
            axios.get('/tasks')
                .then(response => {
                    this.tasks = response.data.items ? response.data.items : [];
                })
                .catch(e => {
                    console.log('Error GET data');
                });
        },
        //Add new task
        createTask() {
            //if name is Empty
            if (!this.newTask.name.trim()) {
                console.log('Empty');
                this.$set(this.newTask, 'name', "");
                return;
            } else {
                //Cut spaces from start ans end of string
                this.$set(this.newTask, 'name', this.newTask.name.trim());

                axios.put('/tasks', this.newTask)
                    .then(response => {
                        this.newTask.id = response.created;
                        this.tasks.push(this.newTask);
                        console.log("Task created!");
                    })
                    .then(() => {
                        this.$set(this.newTask, 'name', "");
                        this.$set(this.newTask, 'id', "");
                    })
                    .catch(error => {
                        console.log('Error PUT data');
                    });
            }
        },
        //Delete task from db
        deleteTask(index) {
            axios.delete('/tasks/' + this.tasks[index].id)
                .then(response => {
                    this.tasks.splice(index, 1);
                    console.log("Task id=" + index + " deleted");
                })
                .catch(error => {
                    console.log('Error DELETE data id=' + index);
                });
        }
    },
});