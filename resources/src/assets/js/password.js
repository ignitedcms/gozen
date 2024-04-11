/*
|---------------------------------------------------------------
| Password component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('password', {
  props: ['value','name'],
  template: `
    <div class="form-group">
      <label for="Password" class="dark:text-white">Password</label>
      <div class="small text-muted dark:text-white">password</div>
      <div class="relative">
        <span @click="eyeball">
          <i
            data-feather="eye"
            class="icon-inside cursor-pointer dark:text-white"
          ></i>
        </span>
        <input
          :name="name"
          :type="textType"
          class="
           form-control
           dark:bg-darkest
           dark:text-white
           dark:border-slate-600"

          placeholder="Password"
        >
      </div>
      <div class="small text-danger"></div>
    </div>
  `,
  data() {
    return {
      textType: 'password',
    };
  },
  methods: {
    eyeball() {
      if (this.textType === 'password') {
        this.textType = 'text';
      } else {
        this.textType = 'password';
      }
    }
  }
});

