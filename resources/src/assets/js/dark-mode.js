/*
|---------------------------------------------------------------
| Dark mode  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('dark-mode', {
  props: ['top'],
  template: `
<button
   @click="toggleDarkMode"
   class="fixed
   z-10
   p-1
   right-4
   border
   border-[--gray]
   bg-white
   dark:bg-darkest
   dark:border-slate-600
   rounded-[--big-radius]"
   :style="{top: top}"
>
  <div class="h-a v-a">
     <i data-feather="sun"
         class="h-[15px] w-[15px] text-gray text-dark dark:hidden"
      ></i>
     <i data-feather="moon"
         class="h-[15px] w-[15px] text-gray text-dark hidden dark:block"
      ></i>
  </div>
</button>
   
  `,
  data() {
    return {
      theme: ''
    }
  },
  methods: {
    toggleDarkMode() {
      if (this.theme == "") {
        document
          .getElementsByTagName("html")[0]
          .setAttribute("class", "dark");
        this.theme = "dark";
      }
      else {
        document.getElementsByTagName("html")[0].removeAttribute("class");
        this.theme = "";
      }
    },
  }
});
