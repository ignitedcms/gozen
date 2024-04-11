/*
|---------------------------------------------------------------
| Switch component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('switch-ios', {
  props: ['name', 'value'],
  template: `
  <div>
     <div class=""></div>
     <label class="form-switch">
        <input
         :name="name"
         type="checkbox"
         role="switch"
         :checked="value"
         @change="handleChange"
        />
        <i></i>
        <div class="switch-text select-none dark:text-white">{{message}}</div>
     </label>
  </div>
  `,
  data() {
    return {
      message: 'Yes/No',
      show: false,
    };
  },
  methods: {
    handleChange(event) {
      this.$emit('input', event.target.checked);
    }
  }
});

