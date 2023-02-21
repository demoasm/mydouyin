#! /bin/bash

gnome-terminal -t "user_service.sh" -- bash -c "sh ./start_script/start_user.sh; exec bash"

gnome-terminal -t "video_service.sh" -- bash -c "sh ./start_script/start_video.sh; exec bash"

gnome-terminal -t "relation_service.sh" -- bash -c "sh ./start_script/start_relation.sh; exec bash"

gnome-terminal -t "favorite_service.sh" -- bash -c "sh ./start_script/start_favorite.sh; exec bash"

gnome-terminal -t "comment_service.sh" -- bash -c "sh ./start_script/start_comment.sh; exec bash"

gnome-terminal -t "message_service.sh" -- bash -c "sh ./start_script/start_message.sh; exec bash"

gnome-terminal -t "api_service.sh" -- bash -c "sh ./start_script/start_api.sh; exec bash"


