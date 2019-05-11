<template>
  <v-layout
    row
    wrap
    justify-center
    align-center
  >
    <v-flex xs12>
      <v-text-field v-model="message"></v-text-field>
      <v-btn color="indigo" @click="echo">Echo</v-btn>
    </v-flex>
  </v-layout>
</template>

<script lang="ts">
import { Vue, Component } from 'vue-property-decorator'
import { useStore } from 'vuex-simple'
import { VStore } from '@/store'
import { grpc } from 'grpc-web-client'
import { EchoService } from '~/proto/echo_pb_service'
import {
  EchoRequest,
  EchoResponse
} from '~/proto/echo_pb'

const host = 'http://localhost:8080'

@Component({
  components: {
  }
})
export default class Index extends Vue {
  private store: VStore = useStore(this.$store)
  public searchWord: string = ''
  public message: string = ''

  public async mounted(): Promise<void> {
    await this.store.illust.refresh()
  }

  public async echo(): Promise<void> {
    const response: EchoResponse = await new Promise<EchoResponse>((resolve) => {
      const request = new EchoRequest()
      request.setMessage(this.message)
      grpc.invoke(EchoService.Echo, {
        request,
        host,
        onMessage: (message: EchoResponse) => {
          resolve(message)
        },
        onEnd: (code) => {
        }
      })
    })
    console.log('EchoService.Echo', response.getMessage())
  }
}
</script>
