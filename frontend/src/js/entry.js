var Vue = require("./vue2.2.5.js");
require("../css/main.css");
require("../css/font-icons.css");


new Vue({
    el: `#app`,
    data: {
        title: `hello world!`
    },
    methods: {
        changeTitle: function(event) {
            this.title = event.target.value;
        }
    }
});




