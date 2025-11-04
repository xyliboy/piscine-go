INTERVIEW_ID=$(grep -R "CLUE" mystery | grep -oE 'interview_[0-9]+' | grep -oE '[0-9]+')
echo "$INTERVIEW_ID"
cat "mystery/interviews/interview_$INTERVIEW_ID"  
echo "$MAIN_SUSPECT"