/*
|---------------------------------------------------------------
| Popover  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('popover', {
  props: ['link', 'width'],
  template: `
    <button
      type="button"
      class="relative rm-btn-styles"
      aria-haspopup="dialog"
      :aria-expanded="arr"
      :aria-controls="'popover-' + uniqueId"
      @keyup.escape="escapePressed()"
      @click="tmp"
      v-click-outside="away"
    >
      <span 
       class="
        border-b-2
        border-[--gray]
        dark:text-white
        dark:border-slate-200"
      > 
      {{link}} 
      </span>
      <div
        :id="'popover-' + uniqueId"
         class="
          fade-in-bottom 
          absolute
          bg-white
          w-[200px]
          p-4
          bottom-[40px] 
          left-[50%]
          ml-[-100px]
          rounded-[--small-radius]
          border
          border-[--gray]
          z-10
          shadow-md
          dark:shadow-none
          dark:bg-darkest
          dark:text-white 
          dark:border-slate-600"

        role="dialog"
        v-if="show"
        :style="{ width: width }"
        @click.stop
      >
        <focus-trap :active="show">
          <slot></slot>
        </focus-trap>
      </div>
    </button>
  `,
  data() {
    return {
      message: '',
      show: false,
      arr: 'false',
      uniqueId: Math.random().toString(36).substring(2) // Generate a unique ID

    };
  },
  methods: {
    away() {
      this.show = false;
      this.arr = 'false';
    },
    tmp() {
      this.show = !this.show;
      if (this.arr == 'false') {
        this.arr = 'true';
      } else {
        this.arr = 'false';
      }
    },
    escapePressed() {
      this.show = false;
      this.arr = 'false';
    }
  }
});

