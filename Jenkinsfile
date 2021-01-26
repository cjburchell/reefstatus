pipeline{
    agent any
    environment {
            DOCKER_IMAGE_SERVER = "cjburchell/reefstatus"
            DOCKER_IMAGE_CONTROLLER = "cjburchell/reefstatus-controller"
            DOCKER_TAG = "${env.BRANCH_NAME}-${env.BUILD_NUMBER}"
            PROJECT_PATH = "/code"
    }

    stages{
        stage('Setup') {
            steps {
                script{
                    slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} started"
                }
				
				/* Let's make sure we have the repository cloned to our workspace */
				checkout scm
             }
        }

        stage('Static Analysis') {
            parallel {
                stage('Go Vet') {
                    agent {
                        docker {
                            image 'cjburchell/goci:1.14'
                            args '-v $WORKSPACE:$PROJECT_PATH'
                        }
                    }
                    steps {
                        script{
                            sh """cd ${PROJECT_PATH} && go vet ./..."""

                            def checkVet = scanForIssues tool: [$class: 'GoVet']
                            publishIssues issues:[checkVet]
                        }
                    }
                }

                stage('Go Lint') {
                    agent {
                        docker {
                            image 'cjburchell/goci:1.14'
                            args '-v $WORKSPACE:$PROJECT_PATH'
                        }
                    }
                    steps {
                        script{
                            sh """cd ${PROJECT_PATH} && golint ./... """

                            def checkLint = scanForIssues tool: [$class: 'GoLint']
                            publishIssues issues:[checkLint]
                        }
                    }
                }
            }
        }

        stage('Build') {
            parallel {
                stage('Build Server') {
                    steps {
                        script {
                            def image = docker.build("${DOCKER_IMAGE_SERVER}")
                            image.tag("${DOCKER_TAG}")
                            if( env.BRANCH_NAME == "master") {
                                image.tag("latest")
                            }
                        }
                    }
                }
                stage('Build Controller') {
                    steps {
                        script {
                            def image = docker.build("${DOCKER_IMAGE_CONTROLLER}", "-f Dockerfile.controller .")
                            image.tag("${DOCKER_TAG}")
                            if( env.BRANCH_NAME == "master") {
                                image.tag("latest")
                            }
                        }
                    }
                }
            }
        }

		stage('Push') {
			parallel {
				stage ('Push Server') {
					steps {
						script {
							docker.withRegistry('', 'dockerhub') {
							   def image = docker.image("${DOCKER_IMAGE_SERVER}")
							   image.push("${DOCKER_TAG}")
							   if( env.BRANCH_NAME == "master") {
									image.push("latest")
							   }
							}
						}
					}
				}
				stage ('Push Controller') {
					steps {
						script {
							docker.withRegistry('', 'dockerhub') {
							   def image = docker.image("${DOCKER_IMAGE_CONTROLLER}")
							   image.push("${DOCKER_TAG}")
							   if( env.BRANCH_NAME == "master") {
									image.push("latest")
							   }
							}
						}
					}
				}
			}
		}
    }

    post {
        always {
              script{
				  sh "docker system prune -f || true"
				  sh "docker image prune -af || true"

                  if ( currentBuild.currentResult == "SUCCESS" ) {
                    slackSend color: "good", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was successful"
                  }
                  else if( currentBuild.currentResult == "FAILURE" ) {
                    slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was failed"
                  }
                  else if( currentBuild.currentResult == "UNSTABLE" ) {
                    slackSend color: "warning", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} was unstable"
                  }
                  else {
                    slackSend color: "danger", message: "Job: ${env.JOB_NAME} with build number ${env.BUILD_NUMBER} its result (${currentBuild.currentResult}) was unclear"
                  }
              }
        }
    }
}