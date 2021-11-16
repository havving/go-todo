
(function($) {
    'use strict';
    $(function() {
        const todoListItem = $('.todo-list');             // 항목이 위치하는 곳
        const todoListInput = $('.todo-list-input');      // 할 일 제목 입력 박스

        $('.todo-list-add-btn').on("click", function(event) {
            event.preventDefault();                     // add 버튼 클릭시

            const item = $(this).prevAll('.todo-list-input').val();

            if (item) {
                $.post("/api/todos", JSON.stringify({name:item}), addItem); // POST 요청
                todoListInput.val("");
            }
        });

        const addItem = function(item) {        // 항목 추가 함수
            if (item.completed) {
                todoListItem.append("<li class='completed'"+ " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' checked='checked' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            } else {
                todoListItem.append("<li "+ " id='" + item.id + "'><div class='form-check'><label class='form-check-label'><input class='checkbox' type='checkbox' />" + item.name + "<i class='input-helper'></i></label></div><i class='remove mdi mdi-close-circle-outline'></i></li>");
            }
        };

        $.get('/api/todos', function(items) {     // 웹 페이지가 뜰때 GET 요청 전송
            items.forEach(e => {
                addItem(e)
            });
        });

        todoListItem.on('change', '.checkbox', function() {
            // 완료 체크 박스 클릭시
            const id = parseInt($(this).closest("li").attr('id'));
            const name = $(this).closest("li").text();
            const $self = $(this);
            let complete = true;
            if ($(this).attr('checked')) {
                complete = false;
            }
            // todos/id로 PUT 요청 전송
            $.ajax({
                url: "/api/todos/" + id,
                type: "PUT",
                data: JSON.stringify({id:id, name:name, completed:complete}),
                success: function(data) {
                    if (complete) {
                        $self.attr('checked', 'checked');
                    } else {
                        $self.removeAttr('checked');
                    }

                    $self.closest("li").toggleClass('completed');
                }
            })
        });


        // 삭제 버튼 클릭 시, DELETE 요청 전송
        todoListItem.on('click', '.remove', function() {
            // url: todos/id method: DELETE
            const id = $(this).closest("li").attr('id');
            const $self = $(this);
            $.ajax({
                url: "/api/todos/" + id,
                type: "DELETE",
                success: function(data) {
                    if (data.success) {
                        $self.parent().remove();
                    }
                }
            })
        });

    });
})(jQuery);