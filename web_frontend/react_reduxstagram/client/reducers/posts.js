// A reducer takes in two things:

// 1. the action (info about what heppend)
// 2. copy of current state

function posts(state = [], action) {
  switch(action.type) {
    case 'INCREMENT_LIKES':
      console.log("incrementing likes")
      const i = action.index
      return [
        ...state.slice(0,i),  // Before the one we update
        {...state[i], likes: state[i].likes + 1},
        ...state.slice(i+1)   // After the one we update
      ]

    default:
      return state
  }
}

export default posts