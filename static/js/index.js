function likeHandler(id){    
    let arr = id.split('_');
    let operation = arr[0];
    let postId = arr[1];
    let likeButton = document.getElementById('like_'+postId);   
    let dislikeButton = document.getElementById('dislike_'+postId);
    let likeSrcArr = likeButton.src.split('/');
    let likeSrc = likeSrcArr[likeSrcArr.length-1];    
    let host = likeButton.src.replace(likeSrc, "");  
    let dislikeSrcArr = dislikeButton.src.split('/');
    let dislikeSrc = dislikeSrcArr[dislikeSrcArr.length-1];
    let likeCounter =  likeButton.previousElementSibling;
    let dislikeCounter =  dislikeButton.previousElementSibling;
    if(operation=="like"){              
        if(dislikeSrc == "dislike_active_icon.png"){
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) - 1;           
        }
        dislikeButton.src = host + "/dislike_inactive_icon.png"      
        if(likeSrc=="like_inactive_icon.png"){
            likeButton.src = host + "/like_active_icon.png";
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) + 1;
            postLikeInfo(postId, 1);
        }else{
            likeButton.src = host + "/like_inactive_icon.png";
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) - 1;
            postLikeInfo(postId, 0);
        }        
    }
    if(operation=="dislike"){       
        if(likeSrc == "like_active_icon.png"){
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) - 1;
        }
        likeButton.src = host + "/like_inactive_icon.png";       
        if(dislikeSrc=="dislike_inactive_icon.png"){
            dislikeButton.src = host + "/dislike_active_icon.png";
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) + 1;
            postLikeInfo(postId, -1);
        }else{
            dislikeButton.src = host + "/dislike_inactive_icon.png";
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) - 1;
            postLikeInfo(postId, 0);
        }        
    }
}
function postLikeInfo(postId, status){
    const endpoint = "/registerlike";
    fetch(endpoint, {
        method: 'POST',
        headers: {'Content-Type':'application/x-www-form-urlencoded'},
        body: "postId="+postId+"&status="+status
    });
}
function commentLikeHandler(id){
    let arr = id.split('_');
    let operation = arr[0].split('-')[1];
    let likeId = arr[1];
    let likeButton = document.getElementById('comment-like_'+likeId);
    let dislikeButton = document.getElementById('comment-dislike_'+likeId);
    let likeSrcArr = likeButton.src.split('/');
    let likeSrc = likeSrcArr[likeSrcArr.length-1];
    
    
    let dislikeSrcArr = dislikeButton.src.split('/');
    let dislikeSrc = dislikeSrcArr[dislikeSrcArr.length-1];
    let likeCounter =  likeButton.previousElementSibling;
    let dislikeCounter =  dislikeButton.previousElementSibling;
   
    if(operation=="like"){              
        if(dislikeSrc == "dislike_active_icon.png"){
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) - 1;           
        }
        dislikeButton.src = "/static/images/dislike_inactive_icon.png"      
        if(likeSrc=="like_inactive_icon.png"){
            likeButton.src = "/static/images/like_active_icon.png";
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) + 1;
            commentLikeInfo(likeId, 1);         
        }else{
            likeButton.src = "/static/images/like_inactive_icon.png";
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) - 1;
            commentLikeInfo(likeId, 0);                   
        }        
    }
    if(operation=="dislike"){       
        if(likeSrc == "like_active_icon.png"){
            likeCounter.innerHTML = parseInt(likeCounter.innerHTML) - 1;
        }
        likeButton.src = "/static/images/like_inactive_icon.png";       
        if(dislikeSrc=="dislike_inactive_icon.png"){
            dislikeButton.src = "/static/images/dislike_active_icon.png";
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) + 1;
            commentLikeInfo(likeId, -1);                  
        }else{
            dislikeButton.src = "/static/images/dislike_inactive_icon.png";
            dislikeCounter.innerHTML = parseInt(dislikeCounter.innerHTML) - 1;
            commentLikeInfo(likeId, 0);                  
        }        
    }
}
function commentLikeInfo(likeId, status){
    const endpoint = "/registercommentlike";
    fetch(endpoint, {
        method: 'POST',
        headers: {'Content-Type':'application/x-www-form-urlencoded'},
        body: "commentId="+likeId+"&status="+status
    });
}
function categoryClickHandler(id){
    let arr = id.split('_');
    let filterOption = arr[1];
    
    let ul = document.getElementById("filter-list");
    
    let elementExists = false;
    for (const child of ul.children) {
        if(child.innerHTML == filterOption){
            elementExists = true;
            break;
        }     
    }    
    if(!elementExists){
        var li = document.createElement("li");
        li.appendChild(document.createTextNode(filterOption));
        li.addEventListener("click", function(e){
            let text = e.target.innerHTML;            
            for (const child of ul.children) {
                if(child.innerHTML == text){
                    ul.removeChild(child);
                    filterContent();                   
                    break;
                }     
            }
        });
        ul.appendChild(li);
        filterContent();
    }    
}
function changeMode(event){
   const endpoint = "/changemode";
   fetch(endpoint, {
       method: 'POST',
       headers: {'Content-Type':'application/x-www-form-urlencoded'},
       body: "mode="+event.target.value
   }).then(() => {
       window.location.href = '/';
   });
}