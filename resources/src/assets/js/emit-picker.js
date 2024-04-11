Vue.component('emitpicker', {
  props: [
    'value'
  ],
  template: `
    <div class="date-picker">
      <input
        type="text"
        :value="value"
        @input="updateDate($event.target.value)"
      />
      <div @click="test('1999-08-06')"> click </div>
    </div>
  `,
  data() {
    return {}
  },
  methods: {
    updateDate(newValue) {
      this.$emit('input', newValue); // Emit the updated value using v-model
    },
    test(idx) {
      this.updateDate(idx);
    },
  },
});

