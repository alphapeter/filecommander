import Vue from 'vue'
import Vuex from 'vuex'
import { Rpc } from '../rpc.js'
import ui from './modules/uiState'

Vue.use(Vuex)

const state = {
  selectedState: 'left',
  roots: [],
  states: {
    left: {
      selectedRoot: '',
      selectedFiles: [],
      path: []
    },
    right: {
      selectedRoot: '',
      selectedFiles: [],
      path: []
    }
  }
}

const actions = {
  init: function ({commit}) {
    Rpc.getRoots()
      .then((response) => {
        commit('init', response.result)
      })
  }
}

const otherStateId = (id) => {
  return id === 'left'
    ? 'right'
    : 'left'
}

const getPathString = path => {
  return path.reduce((acc, p) => {
    return acc + '/' + p
  }, '')
}

const getters = {
  currentState (state) {
    return state.states[state.selectedState]
  },
  currentPathString (state) {
    let path = state.states[state.selectedState].path
    return getPathString(path)
  },
  otherState (state) {
    return state.states[otherStateId(state.selectedState)]
  },
  otherPathString (state) {
    let otherState = state.states[otherStateId(state.selectedState)]
    return getPathString(otherState.path)
  }
}
export const store = new Vuex.Store({
  state: state,
  getters: getters,
  mutations: {
    init (state, roots) {
      state.roots = roots
      state.ui.state = 'browse'
      state.loading = false
      if (roots.length > 0) {
        state.states['left'].selectedRoot = state.roots[0]
        state.states['right'].selectedRoot = roots.length > 1
          ? state.roots[1]
          : state.roots[0]
      }
    },
    selectRoot (state, message) {
      var viewState = state.states[message.stateId]
      viewState.selectedRoot = message.value
      viewState.path = []
    },
    selectSingleFile (state, message) {
      if (state.selectedState !== message.stateId) {
        state.states[state.selectedState].selectedFiles = []
      }
      state.states[message.stateId].selectedFiles = [message.value]
    },
    selectFile (state, message) {
      if (state.selectedState !== message.stateId) {
        state.states[state.selectedState].selectedFiles = []
      }

      let selectedFiles = state.states[message.stateId].selectedFiles
      let fileIndex = selectedFiles.indexOf(message.value)
      if (fileIndex === -1) {
        selectedFiles.push(message.value)
      } else {
        selectedFiles.splice(fileIndex, 1)
      }
    },
    selectFiles (state, message) {
      if (state.selectedState !== message.stateId) {
        state.states[state.selectedState].selectedFiles = []
      }
      state.states[message.stateId].selectedFiles = message.value
    },
    selectView (state, viewId) {
      if (state.selectedState !== viewId) {
        state.states[state.selectedState].selectedFiles = []
      }
      state.selectedState = viewId
    },
    changePath (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.push(message.value)
    },
    changePathToParent (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.pop()
    },
    setPath (state, message) {
      let viewState = state.states[message.stateId]
      viewState.selectedFiles = []
      viewState.path.splice(message.value, viewState.path.length)
    }
  },
  actions: actions,
  modules: {
    ui
  }
})
