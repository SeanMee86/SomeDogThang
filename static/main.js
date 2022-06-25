const btn = document.getElementById('getDogsBtn')
const incBtn = document.getElementById('incBtn')
const decBtn = document.getElementById('decBtn')
const input = document.querySelector('input')
incBtn.addEventListener('click', (e) => {
    input.value = +input.value + 1
})
decBtn.addEventListener('click', (e) => {
    if(+input.value === 1) return;
    input.value = +input.value - 1
})
btn.addEventListener("click", () => {
    document.getElementById("dogImages").innerHTML = ""
    fetch(`http://localhost:8080/api/dogs/${input.value}`)
        .then(resp => resp.json())
        .then(data => {
            data.forEach(dog => {
                const imgUrl = JSON.parse(dog).message
                const imgNode = new Image()
                document.getElementById("dogImages").appendChild(imgNode)
                imgNode.src = imgUrl
            })
        })

})