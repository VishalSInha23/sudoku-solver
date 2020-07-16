function start() {
  localStorage.clear()

  const url = window.location.href

  console.log('Saving service URL.', url)
  setUrl(url)
}

function reset() {
  const url = getUrl()
  window.location.replace(url)
}

async function solve() {
  saveInputSudoku()
  await solveSudoku()
}

async function solveSudoku() {
  const inputSudoku  = getInputSudoku(),
        outputSudoku = await solveSudokuInternal(inputSudoku)

  setOutputSudoku(outputSudoku)

  window.location.href = 'solved.html'
}

function solved() {
  printOutputSudoku()
}

/*------------------------------------------------------------------------------
   UTILITY FUNCTIONS
------------------------------------------------------------------------------*/

async function solveSudokuInternal(sudokuArr) {

  const url    = getUrl() + 'sudoku',
        params = {sudoku : sudokuArr}

  console.log('API Request.', url, params)

  const responseRaw = await fetch(url, {
    method  : 'POST',
    headers : { 'content-type' : 'application/json' },
    body    : JSON.stringify(params)
  })

  const response = await responseRaw.json()

  console.log('API Response.', response)

  return response.sudoku
}

function saveInputSudoku() {
  const arr = []

  for(let i = 0; i < 9; i++) {
    arr[i] = []

    for(let j = 0; j < 9; j++) {
      arr[i][j] = Number(document.getElementById(`cell-${i}${j}`).value) || 0
    }
  }

  console.log('Saving input sudoku.', arr)
  setInputSudoku(arr)
}

function printOutputSudoku() {
  const arr = getOutputSudoku()

  for(let i = 0; i < 9; i++) {
    for(let j = 0; j < 9; j++) {
      document.getElementById(`cell-${i}${j}`).value = arr[i][j]
    }
  }
}

function setUrl(url) {
  localStorage.setItem('url', url)
}

function getUrl() {
  return localStorage.getItem('url')
}

function setInputSudoku(sudoku) {
  localStorage.setItem('input-sudoku', JSON.stringify(sudoku))
}

function getInputSudoku() {
  return JSON.parse(localStorage.getItem('input-sudoku'))
}

function setOutputSudoku(sudoku) {
  localStorage.setItem('output-sudoku', JSON.stringify(sudoku))
}

function getOutputSudoku() {
  return JSON.parse(localStorage.getItem('output-sudoku'))
}