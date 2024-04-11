/*
|---------------------------------------------------------------
| Avatar  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('avatar', {
  props: ['alt','path'],
  template: `
   <img 
    :src="path"
    :alt="alt"
    class="
     shadow-md 
     w-[60px] 
     h-[60px] 
     rounded-full 
     overflow-hidden
     border 
     border-[--gray]
     dark:border-white
    "
   >
   </img>
  `,
  data() {
    return {}
  },
});



