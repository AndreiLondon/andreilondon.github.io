function onInput(){
    var list = document.getElementById("list_categories")
    var optionValue = list.value;
    var categoties = document.getElementById('categories');
    //1 Add element to Selected Categories and Update categoties
    const node = document.createElement("li");
    const textnode = document.createTextNode(optionValue + "   ");
    node.appendChild(textnode);
    document.getElementById("selected_categories").style.display = "inline";
    node.addEventListener("click",function(e){      
        removeElementFromList(optionValue); 
    });
     if (optionValue != "no categoty"){
         document.getElementById("selected_categories").appendChild(node);               
     }
    //2. Remove element from list
    var newOptions = '';
    for (let i = 0; i< list.options.length; i++){
       let ov = list[i].value;
       if (ov != optionValue){
        newOptions+='<option value="' + ov + '" >' + ov + '</option>';
       }      
    }
    list.innerHTML = newOptions;
    
    //3 Update hidden
    var ul = document.getElementById("selected_categories")
    var items = ul.getElementsByTagName("li");
    var freshCategories = '';
    for(let i=0; i<items.length; i++){       
        freshCategories += items[i].textContent;
        if (i < items.length-1){
            freshCategories+=", ";
        }
    }
    categoties.value = freshCategories;    
}
function removeElementFromList(option){   
//1. Remove from selected_categories <ul>
selectedCategories = document.getElementById("selected_categories").getElementsByTagName("li");
    for (let i=0; i< selectedCategories.length; i++){
        val = selectedCategories[i].innerHTML;
        if(val.includes(option)){
            selectedCategories[i].remove();
        }
    }
    if (selectedCategories.length == 0){
        document.getElementById("selected_categories").style.display = "none";
    }
    //2. Add element to list_categories
    var list = document.getElementById("list_categories");
    newOptions = '';
    for (let i = 0; i< list.options.length; i++){
    let ov = list[i].value;  
    newOptions+='<option value="' + ov + '" >' + ov + '</option>';      
    }
    newOptions+='<option value="' + option + '" >' + option + '</option>';
    list.innerHTML = newOptions;
    //3. Update categories (hidden)
    var categoties = document.getElementById('categories');
    categoties.value = '';
    var ul = document.getElementById("selected_categories")
    var items = ul.getElementsByTagName("li");
    var freshCategories = '';
    for(let i=0; i<items.length; i++){
        freshCategories += items[i].textContent;
        if (i < items.length-1){
            freshCategories+=", ";
        }
    }
    categoties.value = freshCategories;
}