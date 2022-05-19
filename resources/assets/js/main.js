
async function searchProduct(customerId) {
    const itemList = getItemList()

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(itemList),
    };
    const url = "/products?search=true&customerId=" + customerId

    await fetch(url, options)
    location.href = url;
}

async function saveOrder(customerId) {

    const finalPrice = document.querySelector("#finalPrice").textContent
    const orderId = document.getElementsByName("id")[0].value

    const itemList = getItemList()

    const customer = {
        id: customerId,
    }
    const order = {
        id: orderId,
        customer: customer,
        items: itemList,
        price: finalPrice,
        time: new Date().toISOString()
    }

    const options = {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(order),
    };

    await fetch("/orders/save", options)
    location.href = "/orders"
}

function checkQuantityUpdatePrice(productStock, index) {
    const list = getItemList()

    let item = list[index]

    let saveButton = document.querySelector(".save-button")
    let span = document.querySelector("#error-"+index)

    if (Number(item.quantity) > Number(productStock) ) {
        span.classList.remove("d-none");
        span.classList.add("animate__headShake");
        saveButton.setAttribute("disabled","")
    } else {
        span.classList.add("d-none");
        span.classList.remove("animate__headShake");
        saveButton.removeAttribute("disabled")
    }

    let finalPrice = 0

    list.forEach( (item) => {
        finalPrice += item.price * item.quantity
    })

    const finalPriceElem = document.querySelector("#finalPrice")
    finalPriceElem.textContent = finalPrice
}

function getItemList() {
    const itemsDivs = document.querySelectorAll(".item-info")

    let list = []
    itemsDivs.forEach( (div) => {

        const quantityInput = div.querySelector("input")
        const quantityValue = quantityInput.value
        const priceValue = div.querySelector(".d-none").textContent
        const productIdValue = quantityInput.getAttribute("id")

        const item = {
            productId: productIdValue,
            price: priceValue,
            quantity: quantityValue
        }
        list.push(item)
    })
    return list

}

function eliminarItem(index) {
    const itemsDivs = document.querySelectorAll(".card")

    itemsDivs[index].remove()
    updateFinalPrice()
}
