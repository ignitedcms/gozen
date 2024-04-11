/*
|---------------------------------------------------------------
| Buttons  component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component("button-component", {
  props: ["variant"],
  template: `
  <button  
     class="
      relative
      inline-block
      text-sm
      px-3
      py-2
      rounded-[--small-radius]"
     :class="[colorVariants[variant]]"
   >
   <slot></slot>
  </button>
  `,
  data() {
    return {
      colorVariants: {
        primary: "bg-[--primary] text-white",
        dark: "bg-dark text-white",
        destructive: "bg-red-600 text-white",
        outline:
          "text-black bg-white  border border-[--gray] hover:bg-gray-200 transition ease-in-out duration-500",
        secondary: "bg-gray-200 text-black ",
      },
    };
  },
});
