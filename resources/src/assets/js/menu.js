/*
|---------------------------------------------------------------
| menu component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('mega-menu', {
  props: [
    'title',
    'logo',
    'url'
  ],
  template: `
    <div 
     class="hide-tablet
      top-0
      sticky
      z-20"
    >
     <div 
      class="
       h-e
       v-a
       bg-white
       p-4
       border-b
       border-[--gray]"
       :aria-labelledby="title"
      >
        <div class="">
          <a
            class="rm-link-styles"
            :href="url"
          >
            <img
              :src="logo"
              class="w-[150px]"
              alt="Ignitedcms logo"
            ></img>
          </a>
        </div>
        <nav class="min-w-[250px]">
          <ul class="rm-list-styles h-e">
          <slot></slot>
          </ul>
        </nav>
        <button-component variant="outline" :id="title">
         {{title}}
        </button-component>
      </div>
    </div>
  `,
  data() {
    return {
      message: ''
    };
  },
  methods: {}
});

Vue.component('menu-items', {
  props: [
    'title',
    'children',
    'url'
  ],
  template: `
    <li @keyup.escape="escapePressed()" class="relative">
      <div
        v-if="children !== 'yes'"
      >
        <a
          :href="url"
          class="rm-link-styles"
        >
          {{title}}
        </a>
      </div>
      <div
        v-if="children === 'yes'"
        class="
         v-a
         relative
         cursor-pointer"

        @click="toggle"
        v-click-outside="away"
      >
        <button
          class="rm-btn-styles"
          :id="title"
          aria-haspopup="true"
          :aria-expanded="show.toString()"
        >
         {{title}}
        </button>
        <span class="ml v-a">
          <i data-feather="chevron-down"></i>
        </span>
      </div>
      <ul
        v-if="show"
        @click.stop
        class="
         absolute 
         top-[40px]
         left-[-10px]
         min-h-[100px]
         min-w-[250px]
         overflow-hidden
         fade-in-bottom
         bg-white
         p
         rounded-[--big-radius]
         border
         border-[--gray]
         shadow-md"
      >
        <slot></slot>
      </ul>
    </li>
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
    },
    escapePressed() {
      this.show = false;
    }
  }
});

Vue.component('menu-item', {
  props: [
    'title',
    'icon',
    'url'
  ],
  template: `
    <li class="row m-t hand">
      <a
        :href="url"
        class="
         rm-link-styles
         rm-list-styles
         p-2
         col
         v-a
         m
         hover:bg-light-gray
         rounded-sm"
      >
        <img
          :src="icon"
          :alt="title"
          class="w-[40px] h-[40px]"
        ></img>
        <div class="ml-2">{{title}}</div>
      </a>
    </li>
  `,
  data() {
    return {};
  },
  methods: {}
});
