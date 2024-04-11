/*
|---------------------------------------------------------------
| Mobile menu component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('mobile-menu', {
  props: [
    'title',
    'logo',
    'url'
  ],
  template: `
    <div 
     class="
     top-0
     sticky
     z-10
     show-tablet"
   >
      <div 
       class="
        h-e
        bg-white
        border-b
        border-light-gray
        p-4
        v-a"
      >
        <div
         class="w-[150px]">
          <a 
            :href="url"
          >
          <img 
            class="v-a"
            :src="logo"
            alt="logo"
          ></img>
          </a>
        </div> 
        <div>
          <span 
            @click="toggle"
          >
            <i 
              data-feather="menu" 
              class="icon cursor-pointer"
            ></i>
          </span>
        </div>
      </div>

      <div 
        v-if="show"
        class="
         fixed
         h-full
         w-full
         border-b
         border-gray-300
         fade-in-bottom
         bg-white" 
      >
        <slot></slot>

        <a 
          href="#" 
          class="rm-link-styles w-[100%]"
        >
          <div 
            class="
             p-4
             bg-white
             border-b
             border-gray-300"
          >
            {{title}}
          </div>
        </a>
      </div>
    </div>
  `,
  data() {
    return {
      show: false
    };
  },
   methods: {
      toggle() {
        this.show = !this.show;
      }
   }
});

Vue.component('mobile-menu-items', {
  props: [
    'title',
    'url',
    'children'
  ],
  template: `
    <div> 
      <a 
        v-if="children !== 'yes'" 
        :href="url" 
        class="rm-link-styles"
      >
        <div 
          class="
           row
           p-4
           bg-white
           v-a 
           border-b
           border-gray-300 
           cursor-pointer"
        >
          {{title}}
        </div>
      </a>

      <div 
        v-if="children === 'yes'" 
        class="
         v-a
         bg-white
         p-4
         h-e
         cursor-pointer
         border-b
         border-gray-300"

        @click="toggle" 
      >
        <div>
          {{title}}
        </div>
        <div>
          +
        </div>
      </div>
      <div 
        v-if="show" 
        class="no-select"
      >
        <slot></slot>
      </div>
    </div>
  `,
  data() {
    return {
      show: false
    };
  },
  methods: {
    toggle() {
      this.show = !this.show;
    },
    away() {
      this.show = false;
    }
  }
});

