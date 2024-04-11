/*
|---------------------------------------------------------------
| Tabs component
|---------------------------------------------------------------
|
|
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/

Vue.component('tabs', {
  template: `
  <div>
    <div 
     class="
      tab-container
      flex
      flex-col
      sm:flex-row
      bg-gray-100
      p-1
      rounded-[--small-radius]
      dark:bg-dark" 

      role="tablist">
      <button
        type="button"
        v-for='(tab, index) in tabs'
        :id="'tab-' +uniqueId+ index"
        role="tab"
        :class='{"bg-white shadow-md dark:bg-darkest dark:text-white": (index == currentIndex)}'
        :aria-selected='index === currentIndex ? "true" : "false"'
        :aria-controls="'tabpanel-'+uniqueId + index"
        :tabindex="currentIndex === index ? 0 : -1"
        class="rm-btn-styles px-2 py-1 rounded-md  dark:text-white"
        @click='selectTab(index)'
        @keydown="onTabKeyDown($event, index)"
        ref="tabButtons"
      >
        {{ tab.title }}
      </button>
   </div>
   <div>
      <slot></slot>
    </div>
  </div>
  `,
  data: function () {
    return {
      currentIndex: 0,
      tabs: [],
      tabIds: [],
      uniqueId: Math.random().toString(36).substring(2) // Generate a unique ID

    }
  },
  created() {
    this.tabs = this.$children
  },
  mounted() {
     // Retrieve the active tab index from localStorage
    const storedIndex = localStorage.getItem('activeTabIndex');
    if (storedIndex !== null) {
      this.currentIndex = parseInt(storedIndex, 10);
    } else {
      this.currentIndex = 0;
    }
    this.selectTab(this.currentIndex);
  },
  methods: {

    selectTab(i) {
        this.currentIndex = i;
      // loop over all the tabs
      this.tabs.forEach((tab, index) => {
        tab.isActive = (index === i);
      });

      // Store the active tab index in localStorage
      localStorage.setItem('activeTabIndex', i);
    },
     onTabKeyDown(event, index) {
        const tabsCount = this.tabs.length;

        switch (event.key) {
           case 'ArrowRight':
              event.preventDefault();
              this.currentIndex = (index + 1) % tabsCount;
              this.$refs.tabButtons[this.currentIndex].focus();
               
              this.selectTab(this.currentIndex);

              break;
           case 'ArrowLeft':
              event.preventDefault();
              this.currentIndex = (index - 1 + tabsCount) % tabsCount;
              this.$refs.tabButtons[this.currentIndex].focus();

              this.selectTab(this.currentIndex);

              break;
           case 'Home':
              event.preventDefault();
              this.currentIndex = 0;
              this.$refs.tabButtons[this.currentIndex].focus();
              break;
           case 'End':
              event.preventDefault();
              this.currentIndex = tabsCount - 1;
              this.$refs.tabButtons[this.currentIndex].focus();
              break;
           case 'Enter':
           case 'Space':
              event.preventDefault();
              this.changeTab(index);
              break;
           default:
              break;
        }
     },

  }
});

Vue.component('tab-item', {
  props: ['title'],
  template: `
    <div
     class=''
     :id="'tabpanel-'+ getIdFromParent+ tabIndex"
     role="tabpanel"
     v-show='isActive'
     :aria-labelledby="'tab-' + getIdFromParent+ tabIndex"
     :aria-hidden="isActive === true ? 'false' : 'true'"
     :tabindex="isActive === true ? 0 : -1"
     ref="tabPanels"
   >
      <slot></slot>
    </div>
  `,
  data: function () {
    return {
      isActive: true,
    }
  },
  computed: {
    tabIndex() {
      // Find the index of the parent tab to associate with aria-labelledby
      return this.$parent.tabs.indexOf(this);
    },
   getIdFromParent() {
        // Access the parent component (tabs)
        const parentTabs = this.$parent;

        // Check if the parent exists and has a uniqueId property
        if (parentTabs && parentTabs.uniqueId) {
          return parentTabs.uniqueId;
        } else {
          //console.error('Parent tabs component not found or missing uniqueId.');
          return null; // or any default value/error handling you prefer
        }
      }
  }
});


