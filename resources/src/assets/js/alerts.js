/*
|---------------------------------------------------------------
| Alert  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('alert', {
  props: ['variant'],
  template: `
    <div
      class="
        bg-white
        p-4
        rounded-[--big-radius]"
        :class="[colorVariants[variant]]"
    >
      <slot></slot>
    </div>
  `,
  data() {
    return {
      colorVariants: {
        success: 'border border-[--gray] dark:bg-dark dark:border-slate-600',
        destructive: 'border border-red-600 text-red-700 dark:bg-red-500',
      },
    };
  },
  methods: {},
});

Vue.component('alert-title', {
  props: ['variant'],
  template: `
    <div class="dark:text-white">
      <slot></slot>
    </div>
  `,
  data() {
    return {};
  },
  methods: {},
});

Vue.component('alert-content', {
  props: ['variant'],
  template: `
    <div class="
      text-sm
      mt-2
      dark:text-white
    ">
      <slot></slot>
    </div>
  `,
  data() {
    return {};
  },
  methods: {},
});

