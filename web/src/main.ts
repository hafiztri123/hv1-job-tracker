import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import './style.css';
import Toast, { PluginOptions, POSITION } from 'vue-toastification'
import 'vue-toastification/dist/index.css'

const options: PluginOptions = {
  hideProgressBar: true,
  position: POSITION.BOTTOM_RIGHT,
  closeOnClick: true,
  closeButton: false,
  timeout: 2000,
}

const app = createApp(App)

app.use(router)
app.use(Toast, options)



app.mount('#app')
