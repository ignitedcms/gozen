/*
|---------------------------------------------------------------
| Toasts component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('toast', {
   props: ['time'],
   template: `
    <div
      v-if="show"
      class="
       toast
       fade-in-bottom
       dark:shadow-none 
       dark:bg-darkest
       dark:border-slate-600"
      role="alert"
      aria-atomic="true"
      aria-live="assertive"
    >
    <slot></slot>
    </div>
  `,
   data() {
      return {
         show: false,
         message: ''
      };
   },
   methods: {
      showToast( duration = 3000) {
         //this.message = message;
         this.show = true;

         setTimeout(() => {
            this.hideToast();
         }, duration);
      },
      hideToast() {
         this.show = false;
      },
   },

});

