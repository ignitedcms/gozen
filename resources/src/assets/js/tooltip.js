/*
|---------------------------------------------------------------
| Tooltip  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('tooltip', {
  props: ['text'],
  template: `
  <div class=" hand" @mouseover="showTooltip" @mouseleave="hideTooltip">
    <slot></slot>
    <div 
     v-if="displayTooltip"
      class="
       relative
       fade-in
       dark:bg-darkest
       dark:shadow-none
       dark:border-slate-600"
    >
      <div
         class="
          small
          fade-in-bottom
          absolute
          bg-white
          w-[200px]
          p-2
          bottom-[40px]
          rounded-[--small-radius]
          border
          border-[--gray]
          text-center
          left-[50%]
          ml-[-100px]
          z-10
          shadow-md
          dark:shadow-none
          dark:bg-darkest
          dark:text-white 
          dark:border-slate-600"
      >

        {{ text }}
      </div>
    </div>
  </div>
  `,
  data() {
    return {
       displayTooltip: false

    };
  },
   methods: {
    showTooltip() {
      this.displayTooltip = true;
    },
    hideTooltip() {
      this.displayTooltip = false;
    }
  }
});


