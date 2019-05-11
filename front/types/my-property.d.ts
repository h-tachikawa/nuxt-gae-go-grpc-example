import { NuxtApp } from '@nuxt/vue-app/types'

declare module '@nuxt/vue-app/types' {
  interface NuxtApp {
    $myProperty: string;
    $testProperty: boolean;
  }
}
