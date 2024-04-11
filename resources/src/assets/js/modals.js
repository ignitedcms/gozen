/*
|---------------------------------------------------------------
| Modals component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('modal', {
  props: [
    'button-title',
    'modal-header'
  ],
  template: `
    <div @keyup.escape="escapePressed()" class="relative inline-block">
      <button
        type="button"
        aria-haspopup="dialog"
        :aria-expanded="arr"
        :aria-controls="'modal-' + uniqueId"
        class=" btn btn-white "
        @click="show=true; arr='true'"
        v-click-outside="away"
      >
        {{buttonTitle}}
      </button>

      <div
        class="
         fixed
         z-20
         w-full
         bg-opacity-80
         h-full
         bg-darker
         left-0
         top-0
         overflow-auto
         v-a
         h-a"

        v-show="show"
        @keyup.escape="escapePressed"
      >
      <div class="show-tablet fixed bottom-0 h-[70%] w-full overflow-hidden rounded-t-lg bg-white p-4 fade-in-bottom" @click.stop>
       <slot></slot>
      </div>

        <div 
          class="
           hide-tablet
           relative
           w-[90%]
           lg:w-[60%]
           shadow-md
           rounded-[--big-radius]
           overflow-hidden
           border
           border-slate-600
           bg-opacity-100
           z-30
           bg-light-gray
           fade-in-bottom
           dark:shadow-none " 

          :id="'modal-' + uniqueId"
          role="dialog"
          @click.stop
        >

          <focus-trap :active="show">
            <div 
             class="
              relative
              bg-white
              border
              border-b-[--gray]
              h-e
              v-a
              px-4
              overflow-hidden
              rounded-t-lg
              dark:border-none
              dark:bg-darker"
            >
              <h5 
               class="
                mt-3
                dark:text-white"
              >
              {{modalHeader}}
              </h5>
              <button
                type="button"
                aria-label="Close"
                class="rm-btn-styles dark:text-white"
                @click="show = false; arr='false'"
              >
              <span>
                 <i data-feather="x"></i>    
              </span>
              </button>
            </div>
            <div class="dark:bg-darker dark:text-white dark:rounded-b-lg">
              <slot></slot>
            </div>
          </focus-trap>

        </div>
      </div>
    </div>
  `,
  data() {
    return {
      message: 'Hello',
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
    escapePressed() {
      this.show = false;
      this.arr = 'false';
    }
  }
});

