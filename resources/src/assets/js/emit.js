Vue.component('emitter',{
    props:['value'],
    template: `
    <input
     class="form-control"
     :value="value"
     @input="$emit('input', $event.target.value)"
    /> 
    `,
    data(){
      return{}
    },
});
