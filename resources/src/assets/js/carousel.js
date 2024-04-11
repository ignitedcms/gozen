/*                                                                          
|---------------------------------------------------------------            
| Carousel component
|---------------------------------------------------------------            
|
| 
| @author: IgnitedCMS
| @license: MIT
| @version: 1.0
| @since: 1.0
|
*/       
Vue.component('carousel', {
  template: `
    <div>
       <div :id="uniqueId" class="splide" aria-labelledby="carousel-heading">
          <div class="splide__track">
             <ul class="splide__list">
                <slot></slot> 
             </ul>
          </div>
       </div>
    </div>
  `,
  data() {
    return {
      uniqueId: 'splide-' + Math.random().toString(36).substring(2) // Generate a unique ID
    }
  },
  mounted() {
    var t = this.uniqueId;
    t = '#'+t.toString();

    new Splide(t, {
       arrows: true,
       pagination:true
    }).mount();
  }
});

Vue.component('carousel-item', {
  props: ['title'],
  template: `
  <li class="splide__slide">
     <div class="
      bg-white
      border 
      border-[--gray]
      shadow-md 
      m-[60px]
      rounded-[--big-radius]
      m-5
      v-a
      h-a
      min-h-[380px]
      dark:bg-darkest
      dark:border-slate-600"
    >
      <h4 class="text-dark">
         <slot></slot>
      </h4>
     </div>
  </li>
  `,
  data() {
    return {

    };
  },
});


