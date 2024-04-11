/*
|---------------------------------------------------------------
| Badges  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('badge', {
  props: ['variant'],
  template: `
  <div
     class="
      inline-block
      select-none
      relative
      text-xs
      px-2
      py-1
      rounded-full"
     :class="[colorVariants[variant]]"
   >
   <slot></slot>
  <div>
  `,
  data() {
     return {
        colorVariants:{
           primary: 'bg-[--primary] text-white',
           dark: 'bg-dark text-white',
           destructive:'bg-red-600 text-white',
           outline:' bg-white text-black border border-black ',
           secondary:'bg-gray-200 text-black '
        }
     }
  },
});
