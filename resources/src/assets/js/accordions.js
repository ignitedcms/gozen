/*                                                                          
|---------------------------------------------------------------            
| Accordion component
|---------------------------------------------------------------            
|
| 
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/       
Vue.component('accordion', {
  template: `
    <div>
      <slot></slot>
    </div>
  `,
  data() {
    return {
      // Nothing here
    };
  },
});

Vue.component('accordion-item', {
  props: ['title'],
  template: `
    <div class="row">
      <div class="col">
        <button
          type="button"
          :aria-expanded="isActive.toString()"
          :aria-controls="'accordion-' + uniqueId"
          :id="'accordion-title-' + uniqueId"
          class="
            h-e
            v-a
            w-full
            border-b
            border-[--gray]
            pb-4
            hover:underline
          "
          @click="toggle"
        >
          <div class="text-black dark:text-white">
            {{ title }}
          </div>
          <span>
            <i data-feather="chevron-down" class="dark:text-white"></i>
          </span>
        </button>
        <div
          v-if="isActive"
          :id="'accordion-' + uniqueId"
          role="region"
          class="
            p-2
            fade-in
          "
          :aria-labelledby="'accordion-title-' + uniqueId"
        >
          <slot ></slot>
        </div>
      </div>
    </div>
  `,
  data() {
    return {
      isActive: false,
      uniqueId: Math.random().toString(36).substring(2), // Generate a unique ID
    };
  },
  methods: {
    toggle() {
      this.isActive = !this.isActive;
    },
  },
});

