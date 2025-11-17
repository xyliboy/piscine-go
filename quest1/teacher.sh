export INTERVIEW_ID=$(head -n 179 streets/Buckingham_Place | tail -n 1 | cut -d'#' -f2)
echo "$INTERVIEW_ID"
cat "./interviews/interview-$INTERVIEW_ID"
echo "$MAIN_SUSPECT"